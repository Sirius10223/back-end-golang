package dtos

import "back-end-golang/helpers"

type UserStatusOKResponse struct {
	StatusCode int                     `json:"status_code" example:"200"`
	Message    string                  `json:"message" example:"Successfully get user credentials"`
	Data       UserInformationResponse `json:"data"`
}

type UserCreeatedResponse struct {
	StatusCode int                     `json:"status_code" example:"201"`
	Message    string                  `json:"message" example:"Successfully registered"`
	Data       UserInformationResponse `json:"data"`
}

type GetAllStationStatusOKResponse struct {
	StatusCode int             `json:"status_code" example:"200"`
	Message    string          `json:"message" example:"Successfully get station"`
	Data       StationResponse `json:"data"`
	Meta       helpers.Meta    `json:"meta"`
}

type StationStatusOKResponse struct {
	StatusCode int             `json:"status_code" example:"200"`
	Message    string          `json:"message" example:"Successfully get station"`
	Data       StationResponse `json:"data"`
}

type StationCreeatedResponse struct {
	StatusCode int             `json:"status_code" example:"201"`
	Message    string          `json:"message" example:"Successfully created station"`
	Data       StationResponse `json:"data"`
}

type GetAllTrainStatusOKResponse struct {
	StatusCode int           `json:"status_code" example:"200"`
	Message    string        `json:"message" example:"Successfully get train"`
	Data       TrainResponse `json:"data"`
	Meta       helpers.Meta  `json:"meta"`
}

type GetAllTrainStatusOKResponses struct {
	StatusCode int            `json:"status_code" example:"200"`
	Message    string         `json:"message" example:"Successfully get train"`
	Data       TrainResponses `json:"data"`
	Meta       helpers.Meta   `json:"meta"`
}

type TrainStatusOKResponse struct {
	StatusCode int           `json:"status_code" example:"200"`
	Message    string        `json:"message" example:"Successfully get train"`
	Data       TrainResponse `json:"data"`
}

type TrainStatusOKResponses struct {
	StatusCode int            `json:"status_code" example:"200"`
	Message    string         `json:"message" example:"Successfully get train"`
	Data       TrainResponses `json:"data"`
}

type TrainCreeatedResponse struct {
	StatusCode int           `json:"status_code" example:"201"`
	Message    string        `json:"message" example:"Successfully created train"`
	Data       TrainResponse `json:"data"`
}

type TrainCreeatedResponses struct {
	StatusCode int            `json:"status_code" example:"201"`
	Message    string         `json:"message" example:"Successfully created train"`
	Data       TrainResponses `json:"data"`
}

type GetAllTrainCarriageStatusOKResponse struct {
	StatusCode int                   `json:"status_code" example:"200"`
	Message    string                `json:"message" example:"Successfully get train peron"`
	Data       TrainCarriageResponse `json:"data"`
	Meta       helpers.Meta          `json:"meta"`
}

type TrainCarriageStatusOKResponse struct {
	StatusCode int                   `json:"status_code" example:"200"`
	Message    string                `json:"message" example:"Successfully get train peron"`
	Data       TrainCarriageResponse `json:"data"`
}

type TrainCarriageCreeatedResponse struct {
	StatusCode int                   `json:"status_code" example:"201"`
	Message    string                `json:"message" example:"Successfully created train peron"`
	Data       TrainCarriageResponse `json:"data"`
}

type GetAllArticleStatusOKResponse struct {
	StatusCode int             `json:"status_code" example:"200"`
	Message    string          `json:"message" example:"Successfully get article"`
	Data       ArticleResponse `json:"data"`
	Meta       helpers.Meta    `json:"meta"`
}

type ArticleStatusOKResponse struct {
	StatusCode int             `json:"status_code" example:"200"`
	Message    string          `json:"message" example:"Successfully get article"`
	Data       ArticleResponse `json:"data"`
}

type ArticleCreeatedResponse struct {
	StatusCode int             `json:"status_code" example:"201"`
	Message    string          `json:"message" example:"Successfully created article"`
	Data       ArticleResponse `json:"data"`
}

type GetAllRecommendationStatusOKResponse struct {
	StatusCode int                    `json:"status_code" example:"200"`
	Message    string                 `json:"message" example:"Successfully get recommendation"`
	Data       RecommendationResponse `json:"data"`
	Meta       helpers.Meta           `json:"meta"`
}

type RecommendationStatusOKResponse struct {
	StatusCode int                    `json:"status_code" example:"200"`
	Message    string                 `json:"message" example:"Successfully get recommendation"`
	Data       RecommendationResponse `json:"data"`
}

type RecommendationCreeatedResponse struct {
	StatusCode int                    `json:"status_code" example:"201"`
	Message    string                 `json:"message" example:"Successfully created recommendation"`
	Data       RecommendationResponse `json:"data"`
}

type HistorySearchStatusOKResponse struct {
	StatusCode int                   `json:"status_code" example:"200"`
	Message    string                `json:"message" example:"Successfully get history search"`
	Data       HistorySearchResponse `json:"data"`
}

type HistorySearchCreeatedResponse struct {
	StatusCode int                   `json:"status_code" example:"201"`
	Message    string                `json:"message" example:"Successfully created history search"`
	Data       HistorySearchResponse `json:"data"`
}

type StatusOKDeletedResponse struct {
	StatusCode int         `json:"status_code" example:"200"`
	Message    string      `json:"message" example:"Successfully deleted"`
	Errors     interface{} `json:"errors"`
}

type BadRequestResponse struct {
	StatusCode int         `json:"status_code" example:"400"`
	Message    string      `json:"message" example:"Bad Request"`
	Errors     interface{} `json:"errors"`
}

type UnauthorizedResponse struct {
	StatusCode int         `json:"status_code" example:"401"`
	Message    string      `json:"message" example:"Unauthorized"`
	Errors     interface{} `json:"errors"`
}

type ForbiddenResponse struct {
	StatusCode int         `json:"status_code" example:"403"`
	Message    string      `json:"message" example:"Forbidden"`
	Errors     interface{} `json:"errors"`
}

type NotFoundResponse struct {
	StatusCode int         `json:"status_code" example:"404"`
	Message    string      `json:"message" example:"Not Found"`
	Errors     interface{} `json:"errors"`
}

type InternalServerErrorResponse struct {
	StatusCode int         `json:"status_code" example:"500"`
	Message    string      `json:"message" example:"Internal Server Error"`
	Errors     interface{} `json:"errors"`
}
