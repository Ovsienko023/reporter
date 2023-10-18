package core

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/app/repository"
	"os"
)

type exportReport struct {
	Id          *string `json:"id,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Date        *int64  `json:"date,omitempty"`
	StartTime   *int64  `json:"start_time,omitempty"`
	EndTime     *int64  `json:"end_time,omitempty"`
	BreakTime   *int64  `json:"break_time,omitempty"`
	WorkTime    *int64  `json:"work_time,omitempty"`
	Body        *string `json:"body,omitempty"`
}

func (c *Core) ExportReportsToJson(ctx context.Context, msg *domain.ExportReportsToJsonRequest) (*domain.ExportReportsToJsonResponse, error) {
	invokerId, err := c.authorize(msg.Token)
	if err != nil {
		return nil, err
	}

	reports, err := c.db.ExportReports(ctx, msg.ToDb(invokerId))
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInternal, err)
	}

	data := struct {
		InvokerId string                      `json:"invoker_id,omitempty"`
		Reports   []repository.ExportedReport `json:"reports,omitempty"`
	}{
		InvokerId: invokerId,
		Reports:   reports,
	}

	jsonData, _ := json.MarshalIndent(data, "", " ")

	fileName := fmt.Sprintf("/tmp/%s.json", invokerId)

	writer, err := os.Create(fmt.Sprintf(fileName))
	if err != nil {
		panic(err.Error())
	}

	defer writer.Close()

	if _, err = writer.Write(jsonData); err != nil {
		panic(err.Error())
	}

	file, err := os.OpenFile(fileName, os.O_RDWR, 0777)
	if err != nil {
		//logger.Error(fmt.Sprintf("Failed to open archive log files: %s", err.Error()))
		return nil, fmt.Errorf("%w: %v", ErrInternal, err)
	}

	return &domain.ExportReportsToJsonResponse{
		//Reports: nil,
		File: file,
	}, nil

	//return new(domain.ExportReportsToJsonResponse).From(result), nil
}
