
export function Auth() {

    this.Email = "";
    this.Password = "";
    this.UserId = 0;

    return this;
}

export function AuthFilter() {

    this.Search = "";
    this.SearchBy = [];
    this.Order = [];
    this.OrderDirection = [];
    this.Ids = [];
    this.CurrentPage = 0;
    this.PerPage = 0;
    this.FunctionType = "";
    this.UrlPath = "";

    return this;
}

export function Authenticator() {

    this.FunctionType = "";
    this.UrlPath = "";

    return this;
}

export function CurrentUser() {

    this.Id = 0;

    return this;
}

export function CurrentUserFilter() {

    this.Search = "";
    this.SearchBy = [];
    this.Order = [];
    this.OrderDirection = [];
    this.Ids = [];
    this.CurrentPage = 0;
    this.PerPage = 0;
    this.FunctionType = "";
    this.UrlPath = "";

    return this;
}

export function Entity() {

    this.Id = 0;
    this.Name = "";
    this.Description = "";

    return this;
}

export function EntityFilter() {

    this.Search = "";
    this.SearchBy = [];
    this.Order = [];
    this.OrderDirection = [];
    this.Ids = [];
    this.CurrentPage = 0;
    this.PerPage = 0;
    this.FunctionType = "";
    this.UrlPath = "";

    return this;
}

export function OrderFilter() {

    this.Order = [];
    this.OrderDirection = [];

    return this;
}

export function Resource() {

    this.Id = 0;
    this.Name = "";
    this.Code = "";
    this.TypeId = 0;

    return this;
}

export function ResourceFilter() {

    this.Search = "";
    this.SearchBy = [];
    this.Order = [];
    this.OrderDirection = [];
    this.Ids = [];
    this.CurrentPage = 0;
    this.PerPage = 0;
    this.FunctionType = "";
    this.UrlPath = "";

    return this;
}

export function ResourceType() {

    this.Id = 0;
    this.Name = "";

    return this;
}

export function ResourceTypeFilter() {

    this.Search = "";
    this.SearchBy = [];
    this.Order = [];
    this.OrderDirection = [];
    this.Ids = [];
    this.CurrentPage = 0;
    this.PerPage = 0;
    this.FunctionType = "";
    this.UrlPath = "";

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

export function Role() {

    this.Id = 0;
    this.Name = "";
    this.Description = "";

    return this;
}

export function RoleFilter() {

    this.Search = "";
    this.SearchBy = [];
    this.Order = [];
    this.OrderDirection = [];
    this.Ids = [];
    this.CurrentPage = 0;
    this.PerPage = 0;
    this.FunctionType = "";
    this.UrlPath = "";

    return this;
}

export function RoleResource() {

    this.Id = 0;
    this.RoleId = 0;
    this.ResourceId = 0;
    this.Find = false;
    this.Read = false;
    this.Create = false;
    this.Update = false;
    this.Delete = false;
    this.FindOrCreate = false;

    return this;
}

export function RoleResourceFilter() {

    this.Search = "";
    this.SearchBy = [];
    this.Order = [];
    this.OrderDirection = [];
    this.Ids = [];
    this.CurrentPage = 0;
    this.PerPage = 0;
    this.FunctionType = "";
    this.UrlPath = "";

    return this;
}

export function User() {

    this.Id = 0;
    this.Email = "";
    this.FirstName = "";
    this.IsActive = false;
    this.LastName = "";
    this.MobilePhone = "";
    this.Password = "";

    return this;
}

export function UserFilter() {

    this.Search = "";
    this.SearchBy = [];
    this.Order = [];
    this.OrderDirection = [];
    this.Ids = [];
    this.CurrentPage = 0;
    this.PerPage = 0;
    this.FunctionType = "";
    this.UrlPath = "";

    return this;
}

export function UserRole() {

    this.Id = 0;
    this.UserId = 0;
    this.RoleId = 0;
    this.RoleID = 0;
    this.UserID = 0;

    return this;
}

export function UserRoleFilter() {

    this.Search = "";
    this.SearchBy = [];
    this.Order = [];
    this.OrderDirection = [];
    this.Ids = [];
    this.CurrentPage = 0;
    this.PerPage = 0;
    this.FunctionType = "";
    this.UrlPath = "";

    return this;
}
