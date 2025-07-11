// Copyright 2025 yumosx
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package event

import "context"

// QueueManager manger for manging the event queue lifecycles per task
type QueueManager interface {
	// Add a new event queue associated with a task Id.
	Add(ctx context.Context, taskId string, queue *Queue) error
	// Get Retrieves the event queue for a task Id.
	Get(ctx context.Context, taskId string) (*Queue, error)
	// Tap Creates a child event queue (tap) for an existing task Id.
	Tap(ctx context.Context, taskId string) (*Queue, error)
	// Close and remove the event queue for a task Id.
	Close(ctx context.Context, taskId string) error
	// CreateOrTap Creates a queue if one doesn't exist, otherwise taps the existing one.
	CreateOrTap(ctx context.Context, taskId string) (*Queue, error)
}
