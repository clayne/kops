// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// Two or more batch entries in the request have the same Id .
type BatchEntryIdsNotDistinct struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *BatchEntryIdsNotDistinct) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BatchEntryIdsNotDistinct) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BatchEntryIdsNotDistinct) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "BatchEntryIdsNotDistinct"
	}
	return *e.ErrorCodeOverride
}
func (e *BatchEntryIdsNotDistinct) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The length of all the messages put together is more than the limit.
type BatchRequestTooLong struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *BatchRequestTooLong) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BatchRequestTooLong) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BatchRequestTooLong) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "BatchRequestTooLong"
	}
	return *e.ErrorCodeOverride
}
func (e *BatchRequestTooLong) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The batch request doesn't contain any entries.
type EmptyBatchRequest struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *EmptyBatchRequest) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EmptyBatchRequest) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EmptyBatchRequest) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "EmptyBatchRequest"
	}
	return *e.ErrorCodeOverride
}
func (e *EmptyBatchRequest) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The accountId is invalid.
type InvalidAddress struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidAddress) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidAddress) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidAddress) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidAddress"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidAddress) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified attribute doesn't exist.
type InvalidAttributeName struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidAttributeName) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidAttributeName) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidAttributeName) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidAttributeName"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidAttributeName) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A queue attribute value is invalid.
type InvalidAttributeValue struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidAttributeValue) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidAttributeValue) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidAttributeValue) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidAttributeValue"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidAttributeValue) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The Id of a batch entry in a batch request doesn't abide by the specification.
type InvalidBatchEntryId struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidBatchEntryId) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidBatchEntryId) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidBatchEntryId) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidBatchEntryId"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidBatchEntryId) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified receipt handle isn't valid for the current version.
type InvalidIdFormat struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidIdFormat) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidIdFormat) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidIdFormat) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidIdFormat"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidIdFormat) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The message contains characters outside the allowed set.
type InvalidMessageContents struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidMessageContents) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidMessageContents) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidMessageContents) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidMessageContents"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidMessageContents) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// When the request to a queue is not HTTPS and SigV4.
type InvalidSecurity struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidSecurity) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSecurity) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSecurity) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidSecurity"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidSecurity) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The caller doesn't have the required KMS access.
type KmsAccessDenied struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *KmsAccessDenied) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KmsAccessDenied) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KmsAccessDenied) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "KmsAccessDenied"
	}
	return *e.ErrorCodeOverride
}
func (e *KmsAccessDenied) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was denied due to request throttling.
type KmsDisabled struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *KmsDisabled) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KmsDisabled) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KmsDisabled) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "KmsDisabled"
	}
	return *e.ErrorCodeOverride
}
func (e *KmsDisabled) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected for one of the following reasons:
//
//   - The KeyUsage value of the KMS key is incompatible with the API operation.
//
//   - The encryption algorithm or signing algorithm specified for the operation
//     is incompatible with the type of key material in the KMS key (KeySpec).
type KmsInvalidKeyUsage struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *KmsInvalidKeyUsage) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KmsInvalidKeyUsage) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KmsInvalidKeyUsage) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "KmsInvalidKeyUsage"
	}
	return *e.ErrorCodeOverride
}
func (e *KmsInvalidKeyUsage) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected because the state of the specified resource is not
// valid for this request.
type KmsInvalidState struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *KmsInvalidState) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KmsInvalidState) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KmsInvalidState) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "KmsInvalidState"
	}
	return *e.ErrorCodeOverride
}
func (e *KmsInvalidState) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected because the specified entity or resource could not be
// found.
type KmsNotFound struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *KmsNotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KmsNotFound) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KmsNotFound) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "KmsNotFound"
	}
	return *e.ErrorCodeOverride
}
func (e *KmsNotFound) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected because the specified key policy isn't syntactically
// or semantically correct.
type KmsOptInRequired struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *KmsOptInRequired) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KmsOptInRequired) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KmsOptInRequired) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "KmsOptInRequired"
	}
	return *e.ErrorCodeOverride
}
func (e *KmsOptInRequired) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Amazon Web Services KMS throttles requests for the following conditions.
type KmsThrottled struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *KmsThrottled) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KmsThrottled) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KmsThrottled) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "KmsThrottled"
	}
	return *e.ErrorCodeOverride
}
func (e *KmsThrottled) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified message isn't in flight.
type MessageNotInflight struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *MessageNotInflight) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MessageNotInflight) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MessageNotInflight) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "MessageNotInflight"
	}
	return *e.ErrorCodeOverride
}
func (e *MessageNotInflight) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified action violates a limit. For example, ReceiveMessage returns this
// error if the maximum number of in flight messages is reached and AddPermission
// returns this error if the maximum number of permissions for the queue is
// reached.
type OverLimit struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *OverLimit) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *OverLimit) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *OverLimit) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "OverLimit"
	}
	return *e.ErrorCodeOverride
}
func (e *OverLimit) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the specified queue previously received a PurgeQueue request
// within the last 60 seconds (the time it can take to delete the messages in the
// queue).
type PurgeQueueInProgress struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *PurgeQueueInProgress) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PurgeQueueInProgress) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PurgeQueueInProgress) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "PurgeQueueInProgress"
	}
	return *e.ErrorCodeOverride
}
func (e *PurgeQueueInProgress) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You must wait 60 seconds after deleting a queue before you can create another
// queue with the same name.
type QueueDeletedRecently struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *QueueDeletedRecently) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *QueueDeletedRecently) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *QueueDeletedRecently) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "QueueDeletedRecently"
	}
	return *e.ErrorCodeOverride
}
func (e *QueueDeletedRecently) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified queue doesn't exist.
type QueueDoesNotExist struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *QueueDoesNotExist) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *QueueDoesNotExist) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *QueueDoesNotExist) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "QueueDoesNotExist"
	}
	return *e.ErrorCodeOverride
}
func (e *QueueDoesNotExist) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A queue with this name already exists. Amazon SQS returns this error only if
// the request includes attributes whose values differ from those of the existing
// queue.
type QueueNameExists struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *QueueNameExists) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *QueueNameExists) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *QueueNameExists) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "QueueNameExists"
	}
	return *e.ErrorCodeOverride
}
func (e *QueueNameExists) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified receipt handle isn't valid.
type ReceiptHandleIsInvalid struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ReceiptHandleIsInvalid) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ReceiptHandleIsInvalid) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ReceiptHandleIsInvalid) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ReceiptHandleIsInvalid"
	}
	return *e.ErrorCodeOverride
}
func (e *ReceiptHandleIsInvalid) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was denied due to request throttling.
//
//   - The rate of requests per second exceeds the Amazon Web Services KMS request
//     quota for an account and Region.
//
//   - A burst or sustained high rate of requests to change the state of the same
//     KMS key. This condition is often known as a "hot key."
//
//   - Requests for operations on KMS keys in a Amazon Web Services CloudHSM key
//     store might be throttled at a lower-than-expected rate when the Amazon Web
//     Services CloudHSM cluster associated with the Amazon Web Services CloudHSM key
//     store is processing numerous commands, including those unrelated to the Amazon
//     Web Services CloudHSM key store.
type RequestThrottled struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *RequestThrottled) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RequestThrottled) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RequestThrottled) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "RequestThrottled"
	}
	return *e.ErrorCodeOverride
}
func (e *RequestThrottled) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// One or more specified resources don't exist.
type ResourceNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ResourceNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The batch request contains more entries than permissible.
type TooManyEntriesInBatchRequest struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *TooManyEntriesInBatchRequest) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyEntriesInBatchRequest) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyEntriesInBatchRequest) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TooManyEntriesInBatchRequest"
	}
	return *e.ErrorCodeOverride
}
func (e *TooManyEntriesInBatchRequest) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Error code 400. Unsupported operation.
type UnsupportedOperation struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *UnsupportedOperation) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedOperation) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedOperation) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "UnsupportedOperation"
	}
	return *e.ErrorCodeOverride
}
func (e *UnsupportedOperation) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }