<?php

namespace App\Http\Controllers;

use App\Http\Requests\BaseRequests\RequestInterface;
use App\Http\Resources\{
    ScheduleResource,
    DefaultCollection,
};;
use App\Services\ScheduleService;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;
use Illuminate\Http\Resources\Json\ResourceCollection;
use Illuminate\Validation\ValidationException;

class ScheduleController extends Controller
{
    public function __construct(ScheduleService $service)
    {
        $this->service = $service;
        $this->jsonResource = ScheduleResource::class;
    }

    public function index(Request $request): ResourceCollection
    {
        $resources = $this->service->index($request->all());

        return new DefaultCollection($this->jsonResource, $resources);
    }

    public function show(int $id): ScheduleResource
    {
        $resource = $this->service->show($id);
        if (!$resource) {
            return $this->notFoundError(['message' => __('messages.api.error.not.found')]);
        }

        return new $this->jsonResource($resource);
    }

    public function store(RequestInterface $request): JsonResponse
    {
        try {
            $data = $request->merge([
                'teacher_id' => auth()->user()->teacher->id,
            ])->all();

            $this->service->store($data);
        } catch (ValidationException $e) {
            return $this->badRequestError(['errors' => $e->errors()]);
        } catch (\Exception $e) {
            return $this->genericError(['message' => $e->getMessage(), 'file' => $e->getFile(), 'line' => $e->getLine()]);
        }

        return $this->noContent();
    }

    public function update(RequestInterface $request, int $id): JsonResponse
    {
        $resource = $this->service->show($id);

        if (!$resource) {
            return $this->_notFoundError(['message' => __('messages.api.error.not.found')]);
        }

        try {
            $updated = $this->service->update($resource, $request->all());
        } catch (ValidationException $e) {
            return $this->badRequestError(['errors' => $e->errors()]);
        } catch (\Exception $e) {
            return $this->genericError(['message' => $e->getMessage(), 'file' => $e->getFile(), 'line' => $e->getLine()]);
        }

        return $this->success([
            'data' => $this->jsonResource::make($updated),
        ]);
    }

    public function destroy(int $id): JsonResponse
    {
        $resource = $this->service->show($id);

        if (!$resource) {
            return $this->notFoundError(['message' => __('messages.api.error.not.found')]);
        }

        try {
            $this->service->destroy($resource);
        } catch (\Exception $e) {
            return $this->genericError(['message' => $e->getMessage(), 'file' => $e->getFile(), 'line' => $e->getLine()]);
        }

        return $this->noContent();
    }
}
