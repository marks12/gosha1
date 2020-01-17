
	//some common actions

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
    this.CurrentId = 0;
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

export function ComponentGroup() {

    this.Id = 0;
    this.Name = "";
    this.Code = "";

    return this;
}

export function ComponentGroupFilter() {

    this.Search = "";
    this.SearchBy = [];
    this.Order = [];
    this.OrderDirection = [];
    this.Ids = [];
    this.CurrentId = 0;
    this.CurrentPage = 0;
    this.PerPage = 0;
    this.FunctionType = "";
    this.UrlPath = "";

    return this;
}

export function ComponentTemplate() {

    this.Id = 0;
    this.Name = "";
    this.Code = "";
    this.Path = "";
    this.GroupCode = "";
    this.GroupId = null;

    return this;
}

export function ComponentTemplateFilter() {

    this.Search = "";
    this.SearchBy = [];
    this.Order = [];
    this.OrderDirection = [];
    this.Ids = [];
    this.CurrentId = 0;
    this.CurrentPage = 0;
    this.PerPage = 0;
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
    this.CurrentId = 0;
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
    this.CurrentId = 0;
    this.CurrentPage = 0;
    this.PerPage = 0;
    this.FunctionType = "";
    this.UrlPath = "";

    return this;
}

export function Layout() {

    this.Id = 0;
    this.Name = "";
    this.RegionId = 0;
    this.LanguageId = 0;

    return this;
}

export function LayoutFilter() {

    this.Search = "";
    this.SearchBy = [];
    this.Order = [];
    this.OrderDirection = [];
    this.Ids = [];
    this.CurrentId = 0;
    this.CurrentPage = 0;
    this.PerPage = 0;
    this.FunctionType = "";
    this.UrlPath = "";

    return this;
}

export function LayoutContent() {

    this.Id = 0;
    this.LayoutId = null;
    this.ComponentTemplateCode = "";
    this.Position = 0;
    this.IsActive = false;
    this.RegionId = null;
    this.LanguageId = null;
    this.Name = "";

    return this;
}

export function LayoutContentFilter() {

    this.Search = "";
    this.SearchBy = [];
    this.Order = [];
    this.OrderDirection = [];
    this.Ids = [];
    this.CurrentId = 0;
    this.CurrentPage = 0;
    this.PerPage = 0;
    this.FunctionType = "";
    this.UrlPath = "";

    return this;
}

export function PageContent() {

    this.Id = 0;
    this.Name = "";
    this.PageTemplateId = null;
    this.IsActive = false;
    this.RegionId = null;
    this.LanguageId = null;
    this.Position = 0;
    this.ComponentTemplateCode = "";

    return this;
}

export function PageContentFilter() {

    this.Search = "";
    this.SearchBy = [];
    this.Order = [];
    this.OrderDirection = [];
    this.Ids = [];
    this.CurrentId = 0;
    this.CurrentPage = 0;
    this.PerPage = 0;
    this.FunctionType = "";
    this.UrlPath = "";

    return this;
}

export function PageInfo() {

    this.Id = 0;
    this.SeoMeta = null;

    return this;
}

export function PageInfoFilter() {

    this.Search = "";
    this.SearchBy = [];
    this.Order = [];
    this.OrderDirection = [];
    this.Ids = [];
    this.CurrentId = 0;
    this.CurrentPage = 0;
    this.PerPage = 0;
    this.FunctionType = "";
    this.UrlPath = "";

    return this;
}

export function PageTemplate() {

    this.Id = 0;
    this.PageTypeId = null;
    this.Name = "";
    this.IsActive = false;
    this.RegionId = null;
    this.LanguageId = null;
    this.RootPageId = null;
    this.LayoutId = null;

    return this;
}

export function PageTemplateFilter() {

    this.Search = "";
    this.SearchBy = [];
    this.Order = [];
    this.OrderDirection = [];
    this.Ids = [];
    this.CurrentId = 0;
    this.CurrentPage = 0;
    this.PerPage = 0;
    this.FunctionType = "";
    this.UrlPath = "";

    return this;
}

export function PageType() {

    this.Id = 0;
    this.Name = "";
    this.Code = "";

    return this;
}

export function PageTypeFilter() {

    this.Search = "";
    this.SearchBy = [];
    this.Order = [];
    this.OrderDirection = [];
    this.Ids = [];
    this.CurrentId = 0;
    this.CurrentPage = 0;
    this.PerPage = 0;
    this.FunctionType = "";
    this.UrlPath = "";

    return this;
}

export function Region() {

    this.Id = 0;
    this.Name = "";
    this.Code = "";
    this.TypeId = null;
    this.EntityId = null;

    return this;
}

export function RegionFilter() {

    this.Search = "";
    this.SearchBy = [];
    this.Order = [];
    this.OrderDirection = [];
    this.Ids = [];
    this.CurrentId = 0;
    this.CurrentPage = 0;
    this.PerPage = 0;
    this.FunctionType = "";
    this.UrlPath = "";

    return this;
}

export function RegionType() {

    this.Id = 0;
    this.Name = "";

    return this;
}

export function RegionTypeFilter() {

    this.Search = "";
    this.SearchBy = [];
    this.Order = [];
    this.OrderDirection = [];
    this.Ids = [];
    this.CurrentId = 0;
    this.CurrentPage = 0;
    this.PerPage = 0;
    this.FunctionType = "";
    this.UrlPath = "";

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
    this.CurrentId = 0;
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
    this.CurrentId = 0;
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
    this.CurrentId = 0;
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
    this.CurrentId = 0;
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
    this.CurrentId = 0;
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

    return this;
}

export function UserRoleFilter() {

    this.Search = "";
    this.SearchBy = [];
    this.Order = [];
    this.OrderDirection = [];
    this.Ids = [];
    this.CurrentId = 0;
    this.CurrentPage = 0;
    this.PerPage = 0;
    this.FunctionType = "";
    this.UrlPath = "";

    return this;
}
