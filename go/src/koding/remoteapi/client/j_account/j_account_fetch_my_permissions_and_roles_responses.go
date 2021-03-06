package j_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// JAccountFetchMyPermissionsAndRolesReader is a Reader for the JAccountFetchMyPermissionsAndRoles structure.
type JAccountFetchMyPermissionsAndRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JAccountFetchMyPermissionsAndRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJAccountFetchMyPermissionsAndRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJAccountFetchMyPermissionsAndRolesOK creates a JAccountFetchMyPermissionsAndRolesOK with default headers values
func NewJAccountFetchMyPermissionsAndRolesOK() *JAccountFetchMyPermissionsAndRolesOK {
	return &JAccountFetchMyPermissionsAndRolesOK{}
}

/*JAccountFetchMyPermissionsAndRolesOK handles this case with default header values.

OK
*/
type JAccountFetchMyPermissionsAndRolesOK struct {
	Payload JAccountFetchMyPermissionsAndRolesOKBody
}

func (o *JAccountFetchMyPermissionsAndRolesOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JAccount.fetchMyPermissionsAndRoles/{id}][%d] jAccountFetchMyPermissionsAndRolesOK  %+v", 200, o.Payload)
}

func (o *JAccountFetchMyPermissionsAndRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*JAccountFetchMyPermissionsAndRolesOKBody j account fetch my permissions and roles o k body
swagger:model JAccountFetchMyPermissionsAndRolesOKBody
*/
type JAccountFetchMyPermissionsAndRolesOKBody struct {
	models.JAccount

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *JAccountFetchMyPermissionsAndRolesOKBody) UnmarshalJSON(raw []byte) error {

	var jAccountFetchMyPermissionsAndRolesOKBodyAO0 models.JAccount
	if err := swag.ReadJSON(raw, &jAccountFetchMyPermissionsAndRolesOKBodyAO0); err != nil {
		return err
	}
	o.JAccount = jAccountFetchMyPermissionsAndRolesOKBodyAO0

	var jAccountFetchMyPermissionsAndRolesOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &jAccountFetchMyPermissionsAndRolesOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = jAccountFetchMyPermissionsAndRolesOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o JAccountFetchMyPermissionsAndRolesOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	jAccountFetchMyPermissionsAndRolesOKBodyAO0, err := swag.WriteJSON(o.JAccount)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jAccountFetchMyPermissionsAndRolesOKBodyAO0)

	jAccountFetchMyPermissionsAndRolesOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jAccountFetchMyPermissionsAndRolesOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this j account fetch my permissions and roles o k body
func (o *JAccountFetchMyPermissionsAndRolesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JAccount.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.DefaultResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
