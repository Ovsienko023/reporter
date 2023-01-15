package handlers

import (
	"errors"
	"fmt"
	"github.com/Ovsienko023/reporter/app/core"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/app/transport/http/httperror"
	"github.com/Ovsienko023/reporter/infrastructure/utils/ptr"
	"net/http"
	"os"
	"strconv"
	"time"
)

func ExportReportsToCsv(c *core.Core, w http.ResponseWriter, r *http.Request) {
	errorContainer := httperror.ErrorResponse{}

	query := r.URL.Query()

	message := domain.GetReportsRequest{
		Token: r.Header.Get("Authorization"),
	}

	// todo Вынести проверки на core + добавить http.Message (продумать генерацию docs)

	dateFrom := query.Get("date_from")
	if dateFrom != "" {
		i, err := strconv.ParseInt(dateFrom, 10, 64)
		if err != nil {
			errorContainer.Done(w, http.StatusBadRequest, "Invalid requests")
			return
		}
		tm := time.Unix(i, 0)
		message.DateFrom = &tm
	}

	dateTo := query.Get("date_to")
	if dateTo != "" {
		i, err := strconv.ParseInt(dateTo, 10, 64)
		if err != nil {
			errorContainer.Done(w, http.StatusBadRequest, "Invalid requests")
			return
		}
		tm := time.Unix(i, 0)
		message.DateTo = &tm
	}

	// todo Добавить отдельную core функцию
	message.Page = ptr.Int(1)
	message.PageSize = ptr.Int(1000)

	result, err := c.GetReports(r.Context(), &message)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrUnauthorized):
			errorContainer.Done(w, http.StatusUnauthorized, err.Error())
			return
		}
		errorContainer.Done(w, http.StatusInternalServerError, "internal error")
		return
	}

	fileBytes := createCsvFile(result.Reports)
	err = FileResponse(w, fileBytes)
	if err != nil {
		errorContainer.Done(w, http.StatusInternalServerError, "internal error")
		return
	}
	deleteFile("reports.csv")
}

func createCsvFile(reports []domain.ReportItem) []byte {
	rawData := "Id,DisplayName,Date,CreatorId,CreatedAt,UpdatedAt,StartTime,EndTime,BreakTime,WorkTime,Body\n"
	for _, report := range reports {

		date := time.Unix(*report.Date, 0)
		createdAt := time.Unix(*report.CreatedAt, 0)
		updatedAt := time.Unix(*report.UpdatedAt, 0)

		rawData += fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s\n",
			strPoint(report.Id),
			strPoint(report.DisplayName),
			date.String(),
			strPoint(report.CreatorId),
			createdAt.String(),
			updatedAt.String(),
			intPoint(report.StartTime),
			intPoint(report.EndTime),
			intPoint(report.BreakTime),
			intPoint(report.WorkTime),
			strPoint(report.Body),
		)

	}
	data := []byte(rawData)
	err := os.WriteFile("reports.csv", data, 0777)
	if err != nil {
		fmt.Println(err)
	}

	return data
}

func deleteFile(filename string) {
	_ = os.Remove(filename)
}

func strPoint(str *string) string {
	if str == nil {
		return ""
	}
	return *str
}

func intPoint(integer *int64) string {
	if integer == nil {
		return "NULL"
	}

	return strconv.Itoa(int(*integer))
}
