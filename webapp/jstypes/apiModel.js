
export function Entity() {

    this.Id = {};
    this.Name = {};
    this.TypeFields = [];
    this.ModelFields = [];

    return this;
}

export function EntityFilter() {

    this.Cats = [];
    this.Search = {};
    this.SearchBy = [];
    this.Order = [];
    this.OrderDirection = [];
    this.CurrentPage = {};
    this.PerPage = {};

    return this;
}

export function EntityField() {

    this.Id = {};
    this.Name = {};
    this.DataType = {};

    return this;
}

export function EntityFieldFilter() {

    this.Cats = [];
    this.Search = {};
    this.SearchBy = [];
    this.Order = [];
    this.OrderDirection = [];
    this.CurrentPage = {};
    this.PerPage = {};

    return this;
}

export function OrderFilter() {

    this.Order = [];
    this.OrderDirection = [];

    return this;
}

export function ProjectInfo() {

    this.Id = {};
    this.Name = {};
    this.Value = {};

    return this;
}

export function ProjectInfoFilter() {

    this.Search = {};
    this.SearchBy = [];
    this.Order = [];
    this.OrderDirection = [];
    this.CurrentPage = {};
    this.PerPage = {};

    return this;
}

export function APIStatus() {

    this.Status = {};

    return this;
}

export function APIError() {

    this.Error = {};
    this.ErrorMessage = {};

    return this;
}

export function Pagination() {

    this.CurrentPage = {};
    this.PerPage = {};

    return this;
}

export function Setting() {

    this.Id = {};

    return this;
}

export function SettingFilter() {

    this.Search = {};
    this.SearchBy = [];
    this.Order = [];
    this.OrderDirection = [];
    this.CurrentPage = {};
    this.PerPage = {};

    return this;
}
