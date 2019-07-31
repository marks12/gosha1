
export function Entity() {

    this.Id = 0;
    this.Desc = "";
    this.Model = "";
    this.Name = "";
    this.VendorID = 0;
    this.URL = "";

    return this;
}

export function EntityFilter() {

    this.Cats = [];
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
