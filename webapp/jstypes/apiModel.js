
	//some common actions

export function BuLayer() {

    this.Id = 0;

    return this;
}

export function BuLayerFilter() {

    this.Search = "";
    this.SearchBy = [];
    this.Order = [];
    this.OrderDirection = [];
    this.CurrentPage = 0;
    this.PerPage = 0;

    return this;
}

export function CurrentApp() {

    this.Id = 0;
    this.IsValidStructure = false;
    this.AdminEmail = "";
    this.AdminPassword = "";
    this.DbType = "";
    this.IsUuidMode = false;
    this.IsViewMode = false;
    this.IsReadonlyMode = false;
    this.IsSoftDelete = false;

    return this;
}

export function CurrentAppFilter() {

    this.Search = "";
    this.SearchBy = [];
    this.Order = [];
    this.OrderDirection = [];
    this.CurrentPage = 0;
    this.PerPage = 0;

    return this;
}

export function Field() {

    this.Name = "";
    this.Type = "";
    this.IsDb = false;
    this.IsType = false;
    this.CommentType = "";
    this.CommentDb = "";

    return this;
}

export function Entity() {

    this.Id = 0;
    this.Name = "";
    this.TypeFields = [];
    this.ModelFields = [];
    this.Fields = [];
    this.IsFilter = false;
    this.CommentType = "";
    this.CommentDb = "";
    this.Structures = {};
    this.HttpMethods = {};
    this.HttpRoutes = {};

    return this;
}

export function EntityFilter() {

    this.Cats = [];
    this.WithFilter = false;
    this.WithHiddenFields = false;
    this.IsFind = false;
    this.IsCreate = false;
    this.IsRead = false;
    this.IsUpdate = false;
    this.IsDelete = false;
    this.IsFindOrCreate = false;
    this.IsUpdateOrCreate = false;
    this.IsRegenerateJsTypes = false;
    this.IsUuidMode = false;
    this.IsViewMode = false;
    this.IsSoftDelete = false;
    this.IsExactMatch = false;
    this.Search = "";
    this.SearchBy = [];
    this.Order = [];
    this.OrderDirection = [];
    this.CurrentPage = 0;
    this.PerPage = 0;

    return this;
}

export function EntityField() {

    this.Id = 0;
    this.Name = "";
    this.DataType = "";

    return this;
}

export function EntityFieldFilter() {

    this.Cats = [];
    this.Search = "";
    this.SearchBy = [];
    this.Order = [];
    this.OrderDirection = [];
    this.CurrentPage = 0;
    this.PerPage = 0;

    return this;
}

export function FieldType() {

    this.Id = 0;
    this.Name = "";
    this.Type = "";

    return this;
}

export function FieldTypeFilter() {

    this.Search = "";
    this.SearchBy = [];
    this.Order = [];
    this.OrderDirection = [];
    this.CurrentPage = 0;
    this.PerPage = 0;

    return this;
}

export function ProjectInfo() {

    this.Id = 0;
    this.Name = "";
    this.Value = "";

    return this;
}

export function ProjectInfoFilter() {

    this.Search = "";
    this.SearchBy = [];
    this.Order = [];
    this.OrderDirection = [];
    this.CurrentPage = 0;
    this.PerPage = 0;

    return this;
}

export function APIStatus() {

    this.Status = "";

    return this;
}

export function APIError() {

    this.Error = false;
    this.ErrorMessage = "";

    return this;
}

export function Pagination() {

    this.CurrentPage = 0;
    this.PerPage = 0;

    return this;
}

export function Setting() {

    this.Id = 0;

    return this;
}

export function SettingFilter() {

    this.Search = "";
    this.SearchBy = [];
    this.Order = [];
    this.OrderDirection = [];
    this.CurrentPage = 0;
    this.PerPage = 0;

    return this;
}
