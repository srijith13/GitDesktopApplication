package helper

// Response is used for static shape json return
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

// BuildResponse method is to inject data value to dynamic success response
func BuildResponse(message string, data interface{}) Response {
	res := Response{
		Status:  true,
		Message: message,
		Error:   nil,
		Data:    data,
	}
	return res
}

// BuildErrorResponse method is to inject data value to dynamic failed response
func BuildErrorResponse(message string, err interface{}, data interface{}) Response {
	res := Response{
		Status:  false,
		Message: message,
		Error:   err,
		Data:    data,
	}
	return res
}

// func ValidateRequestBody(request dto.RouteRequest, params string) []string {
// 	prms := strings.Split(params, ",")
// 	var errors []string
// 	for _, item := range prms {
// 		if item == "DepartureTime" && request.DepartureTime == "" {
// 			errors = append(errors, "departureTime key is required")
// 		} else if item == "SourceLat" && request.SourceLat == 0 {
// 			errors = append(errors, "sourceLat key is required")
// 		} else if item == "SourceLong" && request.SourceLong == 0 {
// 			errors = append(errors, "sourceLong key is required")
// 		} else if item == "DestinationLat" && request.DestinationLat == 0 {
// 			errors = append(errors, "destinationLat key is required")
// 		} else if item == "DestinationLong" && request.DestinationLong == 0 {
// 			errors = append(errors, "destinationLong key is required")
// 		} else if item == "Direction" && request.Direction == "" {
// 			errors = append(errors, "direction key is required")
// 		} else if item == "StartStopId" && request.StartStopId == -1 {
// 			errors = append(errors, "startStopId key is required")
// 		} else if item == "EndStopId" && request.EndStopId == -1 {
// 			errors = append(errors, "endStopId key is required")
// 		} else if item == "RouteId" && request.RouteId == "" {
// 			errors = append(errors, "routeId key is required")
// 		} else if item == "Service" && request.Service == "" {
// 			errors = append(errors, "service key is required")
// 		} else if item == "Search" && request.Search == "" {
// 			errors = append(errors, "search key is required")
// 		} else if item == "Radius" && request.Radius == 0 {
// 			errors = append(errors, "radius key is required")
// 		} else if item == "Limit" && request.Limit == 0 {
// 			errors = append(errors, "limit key is required")
// 		} else if item == "Page" && request.Page == 0 {
// 			errors = append(errors, "page key is required")
// 		} else if item == "Code" && request.Code == "" {
// 			errors = append(errors, "code key is required")
// 		} else if item == "ConcessionType" && request.ConcessionType == "" {
// 			errors = append(errors, "ConcessionType key is required")
// 		} else if item == "AdultBaseFare" && request.AdultBaseFare == 0 {
// 			errors = append(errors, "AdultBaseFare key is required")
// 		} else if item == "BusType" && request.BusType == nil {
// 			errors = append(errors, "BusType key is required")
// 		}
// 	}
// 	validateRegex(request, prms, &errors)
// 	return errors
// }

// func ValidateCustomBody(request dto.RouteRequest) []string {
// 	var errors []string
// 	if ((request.SourceLat == 0 || request.SourceLong == 0 ||
// 		request.DestinationLat == 0 || request.DestinationLong == 0) &&
// 		(request.StartStopId == -1 || request.EndStopId == -1)) || ((request.SourceLat == 0 || request.SourceLong == 0 ||
// 		request.DestinationLat == 0 || request.DestinationLong == 0) &&
// 		(request.StartStopId == 0 && request.EndStopId == 0)) {
// 		errors = append(errors, "Either pass sourceLat, sourceLong, destinationLat, destinationLong or startStopId, endStopId")
// 	}
// 	return errors
// }

// func validateRegex(request dto.RouteRequest, prms []string, errors *[]string) {
// for _, item := range prms {
// 	if item == "DepartureTime" && !regexp.MustCompile(`^([0-9]|0[0-9]|1[0-9]|2[0-3]):([0-9]|[0-5][0-9]):([0-9]|[0-5][0-9])$`).MatchString(request.DepartureTime) {
// 		*errors = append(*errors, "departureTime key is invalid format")
// 	} else if item == "Direction" && !regexp.MustCompile(`^[A-Za-z]*$`).MatchString(request.Direction) {
// 		*errors = append(*errors, "direction key is invalid format")
// 	} else if item == "RouteId" && !regexp.MustCompile(`^[A-Za-z0-9_-]*$`).MatchString(request.RouteId) {
// 		*errors = append(*errors, "routeId key is invalid format")
// 	} else if item == "Service" && !regexp.MustCompile(`^[A-Za-z]*$`).MatchString(request.Service) {
// 		*errors = append(*errors, "service key is invalid format")
// 	} else if item == "Search" && !regexp.MustCompile(`^[\sA-Za-z0-9._-]*$`).MatchString(request.Search) {
// 		*errors = append(*errors, "search key is invalid format")
// 	} else if item == "Code" && !regexp.MustCompile(`^[0-9]*$`).MatchString(request.Code) {
// 		*errors = append(*errors, "code key is invalid format")
// 	}
// }
// }
