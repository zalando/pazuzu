package features

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"swaggen/models"
)

// PostAPIFeaturesReader is a Reader for the PostAPIFeatures structure.
type PostAPIFeaturesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPIFeaturesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostAPIFeaturesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostAPIFeaturesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostAPIFeaturesCreated creates a PostAPIFeaturesCreated with default headers values
func NewPostAPIFeaturesCreated() *PostAPIFeaturesCreated {
	return &PostAPIFeaturesCreated{}
}

/*PostAPIFeaturesCreated handles this case with default header values.

Feature successfully created
*/
type PostAPIFeaturesCreated struct {
	Payload *models.Feature
}

func (o *PostAPIFeaturesCreated) Error() string {
	return fmt.Sprintf("[POST /api/features][%d] postApiFeaturesCreated  %+v", 201, o.Payload)
}

func (o *PostAPIFeaturesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Feature)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIFeaturesDefault creates a PostAPIFeaturesDefault with default headers values
func NewPostAPIFeaturesDefault(code int) *PostAPIFeaturesDefault {
	return &PostAPIFeaturesDefault{
		_statusCode: code,
	}
}

/*PostAPIFeaturesDefault handles this case with default header values.

Unexpected error
*/
type PostAPIFeaturesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post API features default response
func (o *PostAPIFeaturesDefault) Code() int {
	return o._statusCode
}

func (o *PostAPIFeaturesDefault) Error() string {
	return fmt.Sprintf("[POST /api/features][%d] PostAPIFeatures default  %+v", o._statusCode, o.Payload)
}

func (o *PostAPIFeaturesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
