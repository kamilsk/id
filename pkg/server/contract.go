package server

import "github.com/kamilsk/passport/pkg/transfer/api/v1/tracker"

// Service defines the behavior of Passport service.
type Service interface {
	// HandleTrackerFingerprintV1 handles an input request.
	HandleTrackerFingerprintV1(tracker.FingerprintRequest) tracker.FingerprintResponse
	// HandleTrackerInstructionV1 handles an input request.
	HandleTrackerInstructionV1(tracker.InstructionRequest) tracker.InstructionResponse
}
