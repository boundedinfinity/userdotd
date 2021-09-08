package system

import (
	"io/fs"

	"github.com/boundedinfinity/userdotd/model"
	"github.com/boundedinfinity/userdotd/pathutil"
)

func (t *System) EmbeddedListAll(request model.EmbeddedListRequest) (model.EmbeddedListResponse, error) {
	pathFilters := []pathutil.PathFilterFunc{
		pathutil.PathFilterDot,
		pathutil.PathFilterEmpty,
		pathutil.PathFilterKeepMarker,
	}

	dirFilters := []pathutil.DirFilterFunc{
		pathutil.DirFilterKeepMarker,
	}

	return t.embeddedList(request, pathFilters, dirFilters)
}

func (t *System) EmbeddedListDirs(request model.EmbeddedListRequest) (model.EmbeddedListResponse, error) {
	pathFilters := []pathutil.PathFilterFunc{
		pathutil.PathFilterDot,
		pathutil.PathFilterEmpty,
		pathutil.PathFilterKeepMarker,
	}

	dirFilters := []pathutil.DirFilterFunc{
		pathutil.DirFilterKeepMarker,
		pathutil.DirFilterFile,
	}

	return t.embeddedList(request, pathFilters, dirFilters)
}

func (t *System) EmbeddedListFiles(request model.EmbeddedListRequest) (model.EmbeddedListResponse, error) {
	pathFilters := []pathutil.PathFilterFunc{
		pathutil.PathFilterDot,
		pathutil.PathFilterEmpty,
		pathutil.PathFilterKeepMarker,
	}

	dirFilters := []pathutil.DirFilterFunc{
		pathutil.DirFilterKeepMarker,
		pathutil.DirFilterDir,
	}

	return t.embeddedList(request, pathFilters, dirFilters)
}

func (t *System) embeddedList(request model.EmbeddedListRequest, pathFilters []pathutil.PathFilterFunc, dirFilters []pathutil.DirFilterFunc) (model.EmbeddedListResponse, error) {
	response := model.EmbeddedListResponse{}

	fn := func(p string, d fs.DirEntry, err error) error {
		response.Paths = append(response.Paths, p)
		return nil
	}

	err := pathutil.WalkDir("embedded://.", fn, pathFilters, dirFilters)

	return response, err
}
