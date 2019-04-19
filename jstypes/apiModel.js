const Auth = {
    Email: '',
    Password: '',
    Token: '',
};

const AuthFilter = {
    model: Auth,
};

const Authenticator = {
    Token: '',
};

const Category = {
    Id: 0,
    Name: '',
    SourceID: 0,
    SourceCatID: 0,
    ParentID: 0,
    SourceParentID: 0,
    SourceParentName: '',
    SourceName: '',
    IsOwn: false,
};

const CategoryFilter = {
    TestFilter: 0,
    model: Category,
};

const Comment = {
    Id: 0,
    AuthorID: 0,
    RecepientID: 0,
    OrderID: 0,
    OrderListID: 0,
};

const CommentFilter = {
    TestFilter: 0,
    model: Comment,
};

const Company = {
    Id: 0,
    Name: '',
    OffName: '',
    Email: '',
};

const CompanyFilter = {
    TestFilter: 0,
    model: Company,
};

const FilterIds = {
    ids: [],
};

const AbstractFilter = {
    request: {},
};

const APIStatus = {
    Status: '',
};

const APIError = {
    Error: false,
    ErrorMessage: '',
};

const Pagination = {
    TotalRecords: 0,
    TotalPages: 0,
    CurrentPage: 0,
    PerPage: 0,
};

const validator = {
    validationErrors: [],
};

