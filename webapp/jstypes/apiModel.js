
export function Entity() {

    this.Id = 0;
    this.Name = "";
    this.TypeFields = [];
    this.ModelFields = [];

    return this;
}

export function EntityFilter() {

    this.Cats = [];
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

export function OrderFilter() {

    this.Order = [];
    this.OrderDirection = [];

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
