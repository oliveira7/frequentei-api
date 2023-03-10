<?php

namespace App\Http\Resources;

use Illuminate\Http\Request;
use Illuminate\Http\Resources\Json\JsonResource;

class EnrollmentResource extends JsonResource
{
    public function toArray(Request $request): array
    {
        return [
            'id' => (int) $this->id,
            'user_id' => (int) $this->user_id,
            'activity_id' => (int) $this->activity_id,
            'status' => (int) $this->status,
            'start_period' => (string) $this->start_period,
            'end_period' => (string) $this->end_period,
            'created_at' => (string) $this->created_at->format('Y-m-d\TH:i:s\Z'),
            'updated_at' => (string) $this->updated_at->format('Y-m-d\TH:i:s\Z'),
        ];
    }
}
