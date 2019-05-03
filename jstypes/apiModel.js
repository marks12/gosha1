
export function Auth() {

    this.Email = "";
    this.Password = "";
    this.Token = "";

    return this;
}

export function AuthFilter() {

    this.CurrentPage = 0;
    this.PerPage = 0;
    this.Token = "";
    this.TotalPages = 0;
    this.TotalRecords = 0;

    return this;
}

export function Authenticator() {

    this.Token = "";

    return this;
}

export function Category() {

    this.Id = 0;
    this.IsOwn = false;
    this.Name = "";
    this.ParentID = 0;
    this.SourceCatID = 0;
    this.SourceID = 0;
    this.SourceName = "";
    this.SourceParentID = 0;
    this.SourceParentName = "";

    return this;
}

export function CategoryFilter() {

    this.CurrentPage = 0;
    this.PerPage = 0;
    this.TestFilter = 0;
    this.Token = "";
    this.TotalPages = 0;
    this.TotalRecords = 0;

    return this;
}

export function Comment() {

    this.AuthorID = 0;
    this.Id = 0;
    this.OrderID = 0;
    this.OrderListID = 0;
    this.RecepientID = 0;

    return this;
}

export function CommentFilter() {

    this.CurrentPage = 0;
    this.PerPage = 0;
    this.TestFilter = 0;
    this.Token = "";
    this.TotalPages = 0;
    this.TotalRecords = 0;

    return this;
}

export function Company() {

    this.Email = "";
    this.Id = 0;
    this.Name = "";
    this.OffName = "";

    return this;
}

export function CompanyFilter() {

    this.CurrentPage = 0;
    this.PerPage = 0;
    this.TestFilter = 0;
    this.Token = "";
    this.TotalPages = 0;
    this.TotalRecords = 0;

    return this;
}

export function CompanyCompanyType() {

    this.CompanyID = 0;
    this.CompanyTypeID = 0;
    this.Id = 0;

    return this;
}

export function CompanyCompanyTypeFilter() {

    this.CurrentPage = 0;
    this.PerPage = 0;
    this.TestFilter = 0;
    this.Token = "";
    this.TotalPages = 0;
    this.TotalRecords = 0;

    return this;
}

export function CompanyType() {

    this.Id = 0;
    this.Name = "";

    return this;
}

export function CompanyTypeFilter() {

    this.CurrentPage = 0;
    this.PerPage = 0;
    this.TestFilter = 0;
    this.Token = "";
    this.TotalPages = 0;
    this.TotalRecords = 0;

    return this;
}

export function CompanyUser() {

    this.CompanyID = 0;
    this.Id = 0;
    this.UserID = 0;

    return this;
}

export function CompanyUserFilter() {

    this.CurrentPage = 0;
    this.PerPage = 0;
    this.TestFilter = 0;
    this.Token = "";
    this.TotalPages = 0;
    this.TotalRecords = 0;

    return this;
}

export function Currency() {

    this.Code = "";
    this.Id = 0;
    this.IsOwn = false;
    this.Name = "";

    return this;
}

export function CurrencyFilter() {

    this.CurrentPage = 0;
    this.PerPage = 0;
    this.TestFilter = 0;
    this.Token = "";
    this.TotalPages = 0;
    this.TotalRecords = 0;

    return this;
}

export function CurrencyRate() {

    this.CurrencyID = 0;
    this.Id = 0;
    this.RateDate = {};
    this.RateToID = 0;
    this.RateValue = 0;

    return this;
}

export function CurrencyRateFilter() {

    this.CurrentPage = 0;
    this.PerPage = 0;
    this.TestFilter = 0;
    this.Token = "";
    this.TotalPages = 0;
    this.TotalRecords = 0;

    return this;
}

export function Entity() {

    this.CreatedAt = {};
    this.DeletedAt = {};
    this.ID = 0;
    this.UpdatedAt = {};

    return this;
}

export function File() {

    this.Id = 0;
    this.URL = "";

    return this;
}

export function FileFilter() {

    this.CurrentPage = 0;
    this.PerPage = 0;
    this.TestFilter = 0;
    this.Token = "";
    this.TotalPages = 0;
    this.TotalRecords = 0;

    return this;
}

export function AbstractFilter() {

    this.CurrentPage = 0;
    this.PerPage = 0;
    this.Token = "";
    this.TotalPages = 0;
    this.TotalRecords = 0;

    return this;
}

export function GitPull() {

    this.Id = 0;
    this.Name = "";

    return this;
}

export function GitPullFilter() {

    this.CurrentPage = 0;
    this.PerPage = 0;
    this.TestFilter = 0;
    this.Token = "";
    this.TotalPages = 0;
    this.TotalRecords = 0;

    return this;
}

export function Order() {

    this.ClientID = 0;
    this.Id = 0;
    this.ManagerID = 0;
    this.StateID = 0;

    return this;
}

export function OrderFilter() {

    this.CurrentPage = 0;
    this.PerPage = 0;
    this.TestFilter = 0;
    this.Token = "";
    this.TotalPages = 0;
    this.TotalRecords = 0;

    return this;
}

export function OrderList() {

    this.CurrencyID = 0;
    this.DelivDate = "";
    this.Id = 0;
    this.OrderID = 0;
    this.PriceBase = 0;
    this.PriceCurrent = 0;
    this.Qty = 0;
    this.SourceID = 0;
    this.UnitID = 0;
    this.UnitStatusID = 0;

    return this;
}

export function OrderListFilter() {

    this.CurrentPage = 0;
    this.PerPage = 0;
    this.TestFilter = 0;
    this.Token = "";
    this.TotalPages = 0;
    this.TotalRecords = 0;

    return this;
}

export function OrderListLog() {

    this.CurrencyID = 0;
    this.DelivDate = "";
    this.Id = 0;
    this.OrderID = 0;
    this.PriceBase = 0;
    this.PriceCurrent = 0;
    this.Qty = 0;
    this.SourceID = 0;
    this.UnitID = 0;
    this.UnitStatusID = 0;

    return this;
}

export function OrderListLogFilter() {

    this.CurrentPage = 0;
    this.PerPage = 0;
    this.TestFilter = 0;
    this.Token = "";
    this.TotalPages = 0;
    this.TotalRecords = 0;

    return this;
}

export function OrderState() {

    this.Id = 0;
    this.Name = "";

    return this;
}

export function OrderStateFilter() {

    this.CurrentPage = 0;
    this.PerPage = 0;
    this.TestFilter = 0;
    this.Token = "";
    this.TotalPages = 0;
    this.TotalRecords = 0;

    return this;
}

export function OrderStateLog() {

    this.Id = 0;
    this.OrderID = 0;
    this.StateID = 0;

    return this;
}

export function OrderStateLogFilter() {

    this.CurrentPage = 0;
    this.PerPage = 0;
    this.TestFilter = 0;
    this.Token = "";
    this.TotalPages = 0;
    this.TotalRecords = 0;

    return this;
}

export function PriceColumn() {

    this.Enabled = false;
    this.Id = 0;
    this.PriceListTypeID = 0;
    this.PriceTypeID = 0;
    this.SourceID = 0;

    return this;
}

export function PriceColumnFilter() {

    this.CurrentPage = 0;
    this.PerPage = 0;
    this.TestFilter = 0;
    this.Token = "";
    this.TotalPages = 0;
    this.TotalRecords = 0;

    return this;
}

export function PriceList() {

    this.CurrencyID = 0;
    this.DelivDate = "";
    this.Enabled = false;
    this.Id = 0;
    this.Price = 0.0;
    this.PriceListTypeID = 0;
    this.PriceTypeID = 0;
    this.SourceID = 0;
    this.UnitID = 0;

    return this;
}

export function PriceListFilter() {

    this.CurrentPage = 0;
    this.PerPage = 0;
    this.TestFilter = 0;
    this.Token = "";
    this.TotalPages = 0;
    this.TotalRecords = 0;

    return this;
}

export function PriceListType() {

    this.Id = 0;
    this.Name = "";
    this.VisibleForClient = false;

    return this;
}

export function PriceListTypeFilter() {

    this.CurrentPage = 0;
    this.PerPage = 0;
    this.TestFilter = 0;
    this.Token = "";
    this.TotalPages = 0;
    this.TotalRecords = 0;

    return this;
}

export function PriceLoader() {

    this.Id = 0;
    this.Name = "";

    return this;
}

export function PriceLoaderFilter() {

    this.CurrentPage = 0;
    this.PerPage = 0;
    this.TestFilter = 0;
    this.Token = "";
    this.TotalPages = 0;
    this.TotalRecords = 0;

    return this;
}

export function PriceRule() {

    this.Id = 0;
    this.PriceColumnID = 0;
    this.PriceFrom = 0;
    this.PriceOperationType = 0;
    this.PriceTo = 0;
    this.ValueAdd = 0;
    this.ValueMult = 0;

    return this;
}

export function PriceRuleFilter() {

    this.CurrentPage = 0;
    this.PerPage = 0;
    this.TestFilter = 0;
    this.Token = "";
    this.TotalPages = 0;
    this.TotalRecords = 0;

    return this;
}

export function PriceType() {

    this.Id = 0;
    this.Name = "";

    return this;
}

export function PriceTypeFilter() {

    this.CurrentPage = 0;
    this.PerPage = 0;
    this.TestFilter = 0;
    this.Token = "";
    this.TotalPages = 0;
    this.TotalRecords = 0;

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
    this.TotalPages = 0;
    this.TotalRecords = 0;

    return this;
}

export function Role() {

    this.Description = "";
    this.Id = 0;
    this.Name = "";

    return this;
}

export function RoleFilter() {

    this.CurrentPage = 0;
    this.PerPage = 0;
    this.TestFilter = 0;
    this.Token = "";
    this.TotalPages = 0;
    this.TotalRecords = 0;

    return this;
}

export function Source() {

    this.Id = 0;
    this.Name = "";

    return this;
}

export function SourceFilter() {

    this.CurrentPage = 0;
    this.PerPage = 0;
    this.TestFilter = 0;
    this.Token = "";
    this.TotalPages = 0;
    this.TotalRecords = 0;

    return this;
}

export function Unit() {

    this.Id = 0;
    this.Model = "";
    this.Name = "";
    this.VendorID = 0;

    return this;
}

export function UnitFilter() {

    this.Cats = [];
    this.CurrentPage = 0;
    this.PerPage = 0;
    this.Token = "";
    this.TotalPages = 0;
    this.TotalRecords = 0;

    return this;
}

export function UnitFile() {

    this.FileID = 0;
    this.Id = 0;
    this.IsMain = false;
    this.UnitID = 0;

    return this;
}

export function UnitFileFilter() {

    this.CurrentPage = 0;
    this.PerPage = 0;
    this.TestFilter = 0;
    this.Token = "";
    this.TotalPages = 0;
    this.TotalRecords = 0;

    return this;
}

export function UnitSource() {

    this.CategoryID = 0;
    this.Id = 0;
    this.SourceID = 0;
    this.SourceUnitID = "";
    this.SourceUnitName = "";
    this.UnitID = 0;

    return this;
}

export function UnitSourceFilter() {

    this.CurrentPage = 0;
    this.PerPage = 0;
    this.TestFilter = 0;
    this.Token = "";
    this.TotalPages = 0;
    this.TotalRecords = 0;

    return this;
}

export function UnitStatus() {

    this.Id = 0;
    this.Name = "";

    return this;
}

export function UnitStatusFilter() {

    this.CurrentPage = 0;
    this.PerPage = 0;
    this.TestFilter = 0;
    this.Token = "";
    this.TotalPages = 0;
    this.TotalRecords = 0;

    return this;
}

export function User() {

    this.Email = "";
    this.FirstName = "";
    this.Id = 0;
    this.IsActive = false;
    this.LastName = "";
    this.MobilePhone = "";
    this.Password = "";
    this.Token = "";

    return this;
}

export function UserFilter() {

    this.CurrentPage = 0;
    this.PerPage = 0;
    this.TestFilter = 0;
    this.Token = "";
    this.TotalPages = 0;
    this.TotalRecords = 0;

    return this;
}

export function UserRole() {

    this.Id = 0;
    this.RoleID = 0;
    this.UserID = 0;

    return this;
}

export function UserRoleFilter() {

    this.CurrentPage = 0;
    this.PerPage = 0;
    this.TestFilter = 0;
    this.Token = "";
    this.TotalPages = 0;
    this.TotalRecords = 0;

    return this;
}

export function Vendor() {

    this.Id = 0;
    this.Name = "";

    return this;
}

export function VendorFilter() {

    this.CurrentPage = 0;
    this.PerPage = 0;
    this.TestFilter = 0;
    this.Token = "";
    this.TotalPages = 0;
    this.TotalRecords = 0;

    return this;
}
