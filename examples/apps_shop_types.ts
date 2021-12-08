// noinspection JSUnusedGlobalSymbols
export interface FCustomComponent {
    id: number,
    code: string,
    content: FPageContent,
    template: {
        path: string
    },
    settings: unknown,
}

// noinspection JSUnusedGlobalSymbols
export interface FPageContent {
    id: number,
    layout: string,
    name: string,
    layoutId?: string,
    position: number,
}

// noinspection JSUnusedGlobalSymbols
export interface FMainMenuItemListTree {
    children: FMainMenuItemListTree[],
    entity: FMenuItem
}

// noinspection JSUnusedGlobalSymbols
export interface FMenu {
    Id: number,
    Name: string,
    RegionId: number,
    LanguageId: number,
    MenuItems: FMenuItem[]
}

// noinspection JSUnusedGlobalSymbols
export interface FMenuItem {
    ColorPrimary: string,
    ColorSecondary: string,
    Description: string,
    FrontendEvents: [],
    GroupNumber: number,
    Id: number,
    Image: {
        Url: string
    },
    ImageId: number, // - не используется
    IsActive: boolean, // не используется
    IsShowInBreadCrumb: boolean,
    IsUndrelined: boolean,
    MenuId: number, // не используется
    MenuItemTypeDesktop: {
        Code: string,
        Id: number, //
        Name: string //
    },
    MenuItemTypeDesktopId: number,
    MenuItemTypeMobile: {
        Code: string,
        Id: number, //
        Name: string //
    },
    MenuItemTypeMobileId: number, //
    Name: string,
    ParentId: number,
    SeoMeta: {
        Chpu: string,
        Description: string, //
        EntityId: number,
        EntityName: string,
        Id: number, //
        IsPermanentRedirect: boolean, //
        Keywords: string, //
        LanguageId: number,
        PageTemplate: {},
        PageTemplateId: number,
        RedirectSeoMeta: {},
        RedirectSeoMetaId: number,
        RegionId: number,
        Title: string
    },
    SeometaId: number,
    Sort: number,
    SuperscriptColor: string,
    SuperscriptText: string,
    Title: string,
    Url: string
}


// noinspection JSUnusedGlobalSymbols
export interface CheckPromocode {
    Id: number,
    Type: string,
    Discount: number
}

// noinspection JSUnusedGlobalSymbols
export interface FrontendEvent {
}


//some common actions

// noinspection JSUnusedGlobalSymbols
export interface AbTest {

    Id: number,
    ParentId: null,
    Parent: {},
    Hypothesis: string,
    A: string,
    B: string,
    Result: string,
    EndCondition: string,
    StatusId: null,
    TypeId: null,
    Status: {},
    Type: {},
    Name: string
}

// noinspection JSUnusedGlobalSymbols
export interface AbTestFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface AbTestStatus {

    Id: number,
    Name: string
}

// noinspection JSUnusedGlobalSymbols
export interface AbTestStatusFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface AbTestType {

    Id: number,
    Name: string
}

// noinspection JSUnusedGlobalSymbols
export interface AbTestTypeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Achievement {

    Id: number,
    Name: string,
    ImageId: number,
    Code: string
}

// noinspection JSUnusedGlobalSymbols
export interface AchievementFilter {

    ImageId: number,
    Code: string,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Action {

    Id: number,
    OraId: number,
    Name: string,
    TechName: string,
    CampaignOraId: number,
    StartDate: null,
    EndDate: null,
    MinCart: number,
    MinCartPoint: number,
    IsDiscounted: false,
    IsMultipleKits: false,
    IsMultiLevel: false,
    IsRepeatable: false,
    IsMargined: false,
    IsAuto: false,
    OriginalPrice: number,
    SetPoint: number,
    SetPrice: number,
    ProductId: number,
    ProductSetId: number,
    VatPercent: number,
    Description: string,
    DescriptionShort: string,
    RegionId: number,
    LanguageId: number,
    CurrencyCode: string,
    CurrencyId: number,
    BackgroundImageId: number,
    BackgroundColor: string,
    ImageId: number,
    ActionProducts: Array<any>,
    Currency: {},
    BackgroundImage: {},
    Image: {},
    SeoMeta: {},
    TypeCode: string,
    ActionPersonalityTypeId: number,
    ActionPersonalityType: {},
    ActionPriceTypeId: number,
    ActionPriceType: {},
    ActionProductSelectionTypeId: number,
    ActionProductSelectionType: {},
    ActionPurshaseTypeId: number,
    ActionPurshaseType: {},
    OchType: number,
    IsPrizeProgram: false
}

// noinspection JSUnusedGlobalSymbols
export interface ActionFilter {

    RegionId: number,
    LanguageId: number,
    TypeId: number,
    IsFilled: false,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ActionFavorite {

    Id: number,
    AuthId: number,
    UserId: number,
    ActionId: number,
    Action: {}
}

// noinspection JSUnusedGlobalSymbols
export interface ActionFavoriteFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ActionPersonalityType {

    Id: number,
    Name: string,
    Description: string,
    Code: string
}

// noinspection JSUnusedGlobalSymbols
export interface ActionPersonalityTypeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ActionPriceType {

    Id: number,
    Name: string,
    Description: string,
    Code: string
}

// noinspection JSUnusedGlobalSymbols
export interface ActionPriceTypeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ActionProduct {

    Id: number,
    ActionId: number,
    ProductId: number,
    ProductGroupId: number,
    CategoryId: number,
    IsOptional: false,
    IsBonus: number,
    Num: number,
    PointBonus: number,
    PriceBonus: number,
    CollectionNum: number,
    CollectionId: number,
    Product: {}
}

// noinspection JSUnusedGlobalSymbols
export interface ActionProductFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ActionProductSelectionType {

    Id: number,
    Name: string,
    Description: string,
    Code: string
}

// noinspection JSUnusedGlobalSymbols
export interface ActionProductSelectionTypeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ActionPurshaseType {

    Id: number,
    Name: string,
    Description: string,
    Code: string
}

// noinspection JSUnusedGlobalSymbols
export interface ActionPurshaseTypeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}


// noinspection JSUnusedGlobalSymbols// noinspection JSUnusedGlobalSymbols
export interface Address {

    Id: number,
    UserId: number,
    CityId: number,
    CountryId: number,
    Name: string,
    Index: string,
    House: string,
    Building: string,
    Apartment: string,
    Mail: string,
    Phone: string,
    IsActive: false,
    IsMain: false,
    StreetTypeId: number,
    StreetType: {},
    Country: {},
    City: {},
    CityName: string,
    CountryCode: string,
    OraStreetType: string,
    OraId: number,
    Street: string,
    Entrance: string,
    Floor: string,
    CourierComment: string
}

// noinspection JSUnusedGlobalSymbols
export interface AddressFilter {

    UserId: number,
    IsActive: false,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface AdminSession {

    Id: number,
    CountryId: null,
    CountryToken: string,
    HomeToken: string,
    SessionId: string
}

// noinspection JSUnusedGlobalSymbols
export interface AdminSessionFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Advantage {

    Id: number,
    RegionId: number,
    LanguageId: number,
    BaseId: number,
    Title: string,
    Text: string,
    ImageUrl: string,
    ComponentTemplateName: string,
    ImageId: number,
    Image: {}
}

// noinspection JSUnusedGlobalSymbols
export interface AdvantageFilter {

    RegionId: number,
    LanguageId: number,
    BaseId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface AppStat {

    Id: number,
    ApplicationId: number,
    UsersQty: number,
    Cpu: number,
    Memory: string,
    IsWorked: false,
    Uptime: number,
    Restart: number,
    Period: number
}

// noinspection JSUnusedGlobalSymbols
export interface AppStatFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Application {

    Id: number,
    DomainName: string,
    RabbitQueueName: string,
    IsActive: false,
    IsHome: false,
    ServerImportPeriod: number,
    LanguageImportPeriod: number,
    DomainImportPeriod: number,
    ProductImportPeriod: number,
    LastUpdate: null,
    Ip: string,
    DefaultLanguageId: number,
    DefaultRegionId: number,
    MigrationId: number,
    BuildVersion: string,
    RegionIds: Array<any>
}

// noinspection JSUnusedGlobalSymbols
export interface ApplicationFilter {

    Domain: string,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ApplicationCountry {

    Id: number,
    RegionId: number,
    ApplicationId: number
}

// noinspection JSUnusedGlobalSymbols
export interface ApplicationCountryFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Area {

    Id: number,
    Name: string,
    DistrictId: number,
    CountryId: number,
    OraId: number
}

// noinspection JSUnusedGlobalSymbols
export interface AreaFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Article {

    Id: number,
    BaseId: number,
    LanguageCode: string,
    DatePublished: null,
    Title: string,
    SubTitle: string,
    Text: string,
    IsPublished: false,
    OraId: number,
    OraBaseId: number
}

// noinspection JSUnusedGlobalSymbols
export interface ArticleFilter {

    LanguageCode: string,
    BaseId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ArticleList {

    Id: number,
    TypeId: number,
    UserId: number,
    Title: string,
    SubTitle: string,
    UserPhoto: {},
    IsPublished: false,
    DatePublished: null,
    Sort: number,
    UserName: string,
    RankCode: number,
    CountryFlagImage: {},
    VideoUrl: string
}

// noinspection JSUnusedGlobalSymbols
export interface ArticleListFilter {

    TypeIds: Array<any>,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ArticleType {

    Id: number,
    BaseId: number,
    Name: string,
    LanguageCode: string,
    OraId: number,
    OraBaseId: number
}

// noinspection JSUnusedGlobalSymbols
export interface ArticleTypeFilter {

    BaseId: number,
    LanguageCode: string,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Auth {

    Id: number,
    Email: string,
    Password: string,
    Contract: string,
    UserId: number,
    User: {},
    IsActive: false,
    Phone: string,
    CityId: number,
    Longitude: number,
    Latitude: number
}

// noinspection JSUnusedGlobalSymbols
export interface AuthFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Banner {

    Id: number,
    RegionId: number,
    LanguageId: number,
    SeoMetaId: number,
    IconImageId: number,
    Text: string,
    IsClosable: false,
    TextAlign: string,
    ImageAlign: string,
    ColorBg: string,
    ColorText: string,
    IsFullWidth: false,
    Url: string,
    IsActive: false,
    DateStart: null,
    DateEnd: null,
    BannerTypeId: number,
    BannerType: {},
    SeoMeta: {},
    IconImage: {},
    Name: string,
    Title: string,
    Label: string,
    Description: string,
    ButtonText: string,
    MobileImageId: number,
    TabletImageId: number,
    DesktopImageId: number,
    MobileImage: {},
    TabletImage: {},
    DesktopImage: {},
    IsOpenInNewPage: false,
    ColorButton: string,
    DateActionStart: null,
    DateActionEnd: null,
    BaseBannerId: number,
    Position: number,
    SubTitle: string,
    BannerDisplayTypeId: number,
    BannerDisplayType: {},
    IsActiveInitial: false,
    ButtonTextSystemTranslateId: number,
    LabelSystemTranslateId: number,
    ButtonTextSystemTranslate: {},
    LabelSystemTranslate: {}
}

// noinspection JSUnusedGlobalSymbols
export interface BannerFilter {

    IsActive: false,
    BaseBannerId: number,
    BannerTypeId: number,
    BannerDisplayTypeId: number,
    DateStart: null,
    DateEnd: null,
    DateNow: null,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface BannerDisplayType {

    Id: number,
    Name: string,
    Code: string
}

// noinspection JSUnusedGlobalSymbols
export interface BannerDisplayTypeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface BannerType {

    Id: number,
    Name: string,
    Code: string,
    Description: string
}

// noinspection JSUnusedGlobalSymbols
export interface BannerTypeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseAdvantage {

    Id: number,
    Name: string
}

// noinspection JSUnusedGlobalSymbols
export interface BaseAdvantageFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseArticle {

    Id: number,
    TypeId: number,
    UserId: number,
    Title: string,
    SubTitle: string,
    Text: string,
    UserPhotoImageId: number,
    OraId: number,
    IsPublished: false,
    Sort: number,
    VideoUrl: string,
    DatePublished: null,
    UserName: string,
    UserCountryId: number,
    UserCountry: {},
    UserPhotoImage: {},
    OraTypeId: number,
    UserPhotoUrl: string,
    UserContract: string,
    UserCountryCode: string
}

// noinspection JSUnusedGlobalSymbols
export interface BaseArticleApi {

    Id: number,
    TypeId: number,
    UserContract: string,
    Title: string,
    SubTitle: string,
    UserPhoto: string,
    OraId: number,
    IsPublished: false,
    Sort: number,
    VideoUrl: string,
    DatePublished: null,
    UserName: string
}

// noinspection JSUnusedGlobalSymbols
export interface BaseArticleFilter {

    UserId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseArticleType {

    Id: number,
    Name: string,
    OraId: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseArticleTypeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseBanner {

    Id: number,
    Name: string
}

// noinspection JSUnusedGlobalSymbols
export interface BaseBannerFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseBlog {

    Id: number,
    BaseBlogAuthorId: number,
    Title: string,
    Alias: string,
    Text: string,
    DateCreated: null,
    DateUpdated: null,
    IsPublished: false,
    OraId: number,
    BaseBlogAuthorOraId: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseBlogFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseBlogAuthor {

    Id: number,
    Name: string,
    Nickname: string,
    ImageId: number,
    OraId: number,
    OraImageUrl: string
}

// noinspection JSUnusedGlobalSymbols
export interface BaseBlogAuthorFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseCategory {

    Id: number,
    OraId: number,
    Name: string,
    ParentId: number,
    Sort: number,
    TypeId: number,
    OraParentId: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseCategoryFilter {

    ParentId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseComposition {

    Id: number,
    Name: string,
    GroupId: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseCompositionFilter {

    GroupId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseDocument {

    Id: number,
    Name: string,
    GroupId: number,
    OraId: string,
    DocumentTypeId: number,
    DocumentType: {},
    Products: Array<any>,
    ProductCodes: Array<any>,
    RegionIds: Array<any>,
    LanguageIds: Array<any>,
    Url: string,
    ExpiresAt: null,
    IsActive: false
}

// noinspection JSUnusedGlobalSymbols
export interface BaseDocumentFilter {

    GroupId: number,
    DocumentTypeId: number,
    File: {},
    Header: {},
    IsExpires: false,
    IsActive: false,
    WithProductData: false,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseFaq {

    Id: number,
    Name: string,
    GroupId: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseFaqFilter {

    GroupId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseIngredient {

    Id: number,
    Name: string,
    BaseIngredientGroupId: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseIngredientFilter {

    BaseIngredientGroupId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseIngredientGroup {

    Id: number,
    Name: string,
    ParentId: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseIngredientGroupFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseLandingCompositionPercent {

    Id: number,
    Name: string
}

// noinspection JSUnusedGlobalSymbols
export interface BaseLandingCompositionPercentFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseLandingCompositionUnit {

    Id: number,
    Name: string
}

// noinspection JSUnusedGlobalSymbols
export interface BaseLandingCompositionUnitFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseMediaPublication {

    Id: number,
    Title: string,
    Description: string,
    ImageId: number,
    Image: {}
}

// noinspection JSUnusedGlobalSymbols
export interface BaseMediaPublicationFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseNews {

    Id: number,
    BaseBlogId: number,
    Title: string,
    Text: string,
    IsPublished: false,
    PreviewImageId: number,
    DateCreated: null,
    DateUpdated: null,
    OraId: number,
    BaseTopicId: number,
    BaseBlogOraId: number,
    PreviewImageUrl: string
}

// noinspection JSUnusedGlobalSymbols
export interface BaseNewsFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseProductIcon {

    Id: number,
    Name: string,
    ImageId: number,
    Image: {},
    AdminName: string,
    IsMarker: false,
    OraId: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseProductIconFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseQualityAssurance {

    Id: number,
    Name: string,
    Html: string
}

// noinspection JSUnusedGlobalSymbols
export interface BaseQualityAssuranceFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseReview {

    Id: number,
    Name: string,
    GroupId: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseReviewFilter {

    GroupId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseStrainer {

    Id: number,
    BaseStrainerTypeId: number,
    Value: string,
    BaseStrainerType: {}
}

// noinspection JSUnusedGlobalSymbols
export interface BaseStrainerFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseStrainerType {

    Id: number,
    Name: string,
    StrainerTypes: Array<any>
}

// noinspection JSUnusedGlobalSymbols
export interface BaseStrainerTypeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseStreetType {

    Id: number,
    Name: string,
    OraId: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseStreetTypeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseSystemTranslate {

    Id: number,
    Name: string,
    Description: string,
    BaseSystemTranslateTypeId: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseSystemTranslateFilter {

    BaseSystemTranslateTypeId: number,
    ExceptRegionId: number,
    ExceptLanguageId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseSystemTranslateType {

    Id: number,
    Name: string
}

// noinspection JSUnusedGlobalSymbols
export interface BaseSystemTranslateTypeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseTopic {

    Id: number,
    Name: string,
    Topic: {}
}

// noinspection JSUnusedGlobalSymbols
export interface BaseTopicFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseTranslate {

    Id: number,
    Keyword: string,
    Translate: string,
    BaseTranslateVars: Array<any>
}

// noinspection JSUnusedGlobalSymbols
export interface BaseTranslateFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface BaseTranslateVar {

    Id: number,
    BaseTranslateCode: string,
    BaseTranslateId: number,
    VarName: string,
    VarDefaultValue: string
}

// noinspection JSUnusedGlobalSymbols
export interface BaseTranslateVarFilter {

    BaseTranslateId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Blog {

    Id: number,
    RegionId: number,
    LanguageId: number,
    BaseBlogId: number,
    Title: string,
    Text: string,
    IsPublished: false,
    OraId: number,
    BaseBlogOraId: number,
    RegionOraCode: string,
    LanguageOraCode: string
}

// noinspection JSUnusedGlobalSymbols
export interface BlogFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface BlogAuthor {

    Id: number,
    Name: string,
    Nickname: string,
    OraId: number,
    RegionId: number,
    LanguageId: number,
    BaseBlogAuthorId: number,
    BaseBlogAuthorOraId: number,
    RegionOraCode: string,
    LanguageOraCode: string
}

// noinspection JSUnusedGlobalSymbols
export interface BlogAuthorFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface BusinessEarningProspect {

    Id: number,
    CurrencyId: number,
    AdditionalIncomeFrom: number,
    AdditionalIncomeTo: number,
    BasicIncomeFrom: number,
    BasicIncomeTo: number,
    BigBusiness: number,
    Currency: {}
}

// noinspection JSUnusedGlobalSymbols
export interface BusinessEarningProspectFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Cart {

    Id: number,
    Code: string,
    AuthId: number,
    CreatedAt: null,
    UpdatedAt: null,
    IsPackageMode: false,
    CartPackages: Array<any>
}

// noinspection JSUnusedGlobalSymbols
export interface CartFilter {

    UserId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface CartChange {

    Id: number,
    ProductId: number,
    OldValue: number,
    NewValue: number,
    ChangeType: string
}

// noinspection JSUnusedGlobalSymbols
export interface CartChangeFilter {

    CartId: number,
    CityId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface CartPackage {

    Id: number,
    CartId: number,
    RecipientUserId: number,
    RecipientUserContract: string,
    IsCurrent: false,
    RecipientAuthId: number,
    RecipientUser: {},
    CartPackageProducts: Array<any>,
    CertificateId: number,
    Certificate: {},
    IsMyPackage: false
}

// noinspection JSUnusedGlobalSymbols
export interface CartPackageFilter {

    CartId: number,
    RecipientUserContract: string,
    RecipientUserId: number,
    RecipientAuthId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface CartPackageProduct {

    Id: number,
    CartPackageId: number,
    ProductId: number,
    HiddenPrice: number,
    Quantity: number
}

// noinspection JSUnusedGlobalSymbols
export interface CartPackageProductFilter {

    CartPackageId: number,
    ProductId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface CartSummary {

    Id: number,
    Cost: number,
    Point: number,
    Discount: number,
    ProductCount: number,
    Currency: {}
}

// noinspection JSUnusedGlobalSymbols
export interface CartSummaryFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface CashBack {

    Id: number,
    CashBackTypeId: number,
    Limit: number,
    Percent: number,
    Total: number,
    CashBackType: {}
}

// noinspection JSUnusedGlobalSymbols
export interface CashBackFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface CashBackRole {

    Id: number,
    RoleId: number,
    CashBackId: number,
    RegionId: number
}

// noinspection JSUnusedGlobalSymbols
export interface CashBackRoleFilter {

    CashBackId: number,
    RoleId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface CashBackType {

    Id: number,
    Name: string
}

// noinspection JSUnusedGlobalSymbols
export interface CashBackTypeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Category {

    Id: number,
    Name: string,
    OrigName: string,
    IsActive: false,
    Description: string,
    ParentId: number,
    TypeId: number,
    Sort: number,
    UrlCode: string,
    ShowInMobile: false,
    RegionId: number,
    LanguageId: number,
    OraId: number,
    BaseId: number,
    ImageId: number,
    SeoDescription: string,
    SeoHeader: string,
    SeoKeywords: string,
    IsDeleted: false,
    OraParentId: number,
    Title: string,
    VideoId: number,
    PromoBackgroundImageId: number,
    PromoTitle: string,
    SeoText: string,
    IsShowInCollectionList: false,
    Url: string,
    SeoMetaId: number,
    Image: {},
    PromoBackgroundImage: {},
    Banners: Array<any>,
    Video: {},
    SeoMeta: {},
    CategoryItems: Array<any>,
    ChildCategories: Array<any>
}

// noinspection JSUnusedGlobalSymbols
export interface CategoryFilter {

    IgnoreSeo: false,
    BaseId: number,
    BaseIds: Array<any>,
    IsPreloadRelation: false,
    IsActive: false,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number,
    OraId: number,
    LanguageId: number,
    RegionId: number,
    ParentId: number,
    IsTop: false,
    TypeId: number
}

// noinspection JSUnusedGlobalSymbols
export interface CategoryBanner {

    Id: number,
    CategoryId: number,
    BannerId: number
}

// noinspection JSUnusedGlobalSymbols
export interface CategoryBannerFilter {

    CategoryId: number,
    BannerId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface CategoryItem {

    Id: number,
    CategoryId: number,
    TypeId: number,
    SourceId: number,
    Position: number,
    IsActive: false,
    ComponentCode: string,
    Type: {},
    IsDefaultCategory: false,
    Volume: number
}

// noinspection JSUnusedGlobalSymbols
export interface CategoryItemFilter {

    CategoryId: number,
    SourceId: number,
    TypeId: number,
    StrainersIds: Array<any>,
    IsNotArchived: false,
    IsActive: false,
    CityId: number,
    IsSiteShow: false,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface CategoryItemType {

    Id: number,
    Name: string
}

// noinspection JSUnusedGlobalSymbols
export interface CategoryItemTypeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface CategoryProductCopy {

    Id: number,
    SourceRegionId: number,
    SourceLanguageId: number,
    DestinationRegionId: number,
    DestinationLanguageId: number,
    AllFoundProductCount: number,
    CopiedProducts: number,
    UserId: number
}

// noinspection JSUnusedGlobalSymbols
export interface CategoryProductCopyFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface CategoryType {

    Id: number,
    Name: string
}

// noinspection JSUnusedGlobalSymbols
export interface CategoryTypeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Certificate {

    Id: number,
    Contract: string,
    CurrencyId: number,
    HeaderOraId: number,
    OraId: number,
    Number: string,
    Stars: number,
    StoreContract: string,
    NewStoreContract: string,
    Total: number,
    Free: number,
    UserId: number,
    Currency: {},
    CertificateTypeId: number,
    CertificateType: {}
}

// noinspection JSUnusedGlobalSymbols
export interface CertificateFilter {

    UserId: number,
    CurrencyId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface CertificateType {

    Id: number,
    Name: string,
    Code: string,
    ImageId: number,
    Image: {}
}

// noinspection JSUnusedGlobalSymbols
export interface CertificateTypeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface CheckPromocodeFilter {

    Promocode: string,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface City {

    Id: number,
    Name: string,
    CountryId: number,
    OraId: number,
    DistrictId: number,
    Lat: number,
    Long: number,
    GoogleGeoId: string,
    TimezoneId: number,
    AreaId: number,
    District: {},
    Area: {},
    IsShowArea: false,
    StoreId: number
}

// noinspection JSUnusedGlobalSymbols
export interface CityFilter {

    CountryId: number,
    OraId: number,
    IsIgnoreGeo: false,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface CityMentor {

    Id: number,
    CityId: number,
    UserId: number,
    Contract: string,
    IsActive: false,
    OraClientServiceCenterContract: string,
    User: {},
    City: {},
    Birthday: null,
    Gender: string
}

// noinspection JSUnusedGlobalSymbols
export interface CityMentorFilter {

    CityId: number,
    WithNoActive: false,
    RegionIds: Array<any>,
    Birthday: null,
    WithStrategy: false,
    ForGender: string,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface CityMentorSearchStrategy {

    Id: number,
    Name: string,
    Description: string,
    IsDefault: false
}

// noinspection JSUnusedGlobalSymbols
export interface CityMentorSearchStrategyFilter {

    IsDefaultOnly: false,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ClientServiceCenter {

    Id: number,
    Name: string,
    OraId: number,
    Latitude: number,
    Longitude: number,
    CountryId: number,
    CountryCode: string,
    TypeId: number,
    PhoneNumber: string,
    Email: string,
    IsCashPayment: false,
    IsCashlessPayment: false,
    IsPayOnDelivery: false,
    CityId: number,
    MinDeliveryCost: number,
    CurrencyId: number,
    StoreContract: string,
    DeliveryPointTypeStr: string,
    Department: string,
    BusinessHours: string,
    CurrencyCode: string,
    Index: string,
    Street: string,
    House: string,
    Apartment: string,
    Building: string,
    CityOraId: number,
    StoreId: number,
    Type: {},
    City: {},
    Country: {},
    Currency: {},
    MetroStations: Array<any>,
    Images: Array<any>,
    HowToGet: string,
    ClientServiceCenterContract: string,
    OraAddressId: number,
    Address: string,
    Sort: number,
    IsFreeDelivery: false
}

// noinspection JSUnusedGlobalSymbols
export interface ClientServiceCenterFilter {

    CountryId: number,
    TypeId: number,
    CityId: number,
    StoreId: number,
    ClientServiceCenterContract: string,
    IsIgnoreType: false,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ClientServiceCenterImage {

    Id: number,
    ClientServiceCenterId: number,
    ImageId: number
}

// noinspection JSUnusedGlobalSymbols
export interface ClientServiceCenterImageFilter {

    ImageId: number,
    ClientServiceCenterId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ClientServiceCenterMetroStation {

    Id: number,
    MetroStationId: number,
    ClientServiceCenterId: number
}

// noinspection JSUnusedGlobalSymbols
export interface ClientServiceCenterMetroStationFilter {

    MetroStationId: number,
    ClientServiceCenterId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ClientServiceCenterType {

    Id: number,
    Name: string,
    Code: string
}

// noinspection JSUnusedGlobalSymbols
export interface ClientServiceCenterTypeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ComplexProgram {

    Id: number,
    Title: string,
    LeftImageId: number,
    RightImageId: number,
    LeftImage: {},
    RightImage: {},
    SeoMetaId: number,
    Description: string,
    RegionId: number,
    LanguageId: number
}

// noinspection JSUnusedGlobalSymbols
export interface ComplexProgramFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ComponentGroup {

    Id: number,
    Name: string,
    Code: string
}

// noinspection JSUnusedGlobalSymbols
export interface ComponentGroupFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ComponentSetting {

    Id: number,
    Name: string,
    Code: string,
    TypeId: number,
    ComponentTemplateCode: string,
    GroupName: string,
    DefaultValue: string,
    IsArray: false
}

// noinspection JSUnusedGlobalSymbols
export interface ComponentSettingFilter {

    ComponentTemplateCode: string,
    TypeId: number,
    Code: string,
    ComponentTemplateCodes: Array<any>,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ComponentTemplate {

    Id: number,
    Name: string,
    Code: string,
    Path: string,
    Url: string,
    GroupCode: string,
    GroupId: number
}

// noinspection JSUnusedGlobalSymbols
export interface ComponentTemplateFilter {

    Codes: Array<any>,
    GroupId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Composition {

    Id: number,
    Name: string,
    Description: string,
    FullDescription: string,
    Value: string,
    RegionId: number,
    LanguageId: number,
    BaseId: number,
    IsActive: false,
    GroupId: number
}

// noinspection JSUnusedGlobalSymbols
export interface CompositionFilter {

    IsActive: false,
    BaseId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface CompositionGroup {

    Id: number,
    Name: string,
    ParentId: number
}

// noinspection JSUnusedGlobalSymbols
export interface CompositionGroupFilter {

    ParentId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ContentSettings {

    Id: number,
    Value: string,
    ComponentSettingId: number,
    ContentId: number,
    ContentType: string,
    ComponentSetting: {},
    IsArray: false,
    ParentId: number,
    Sort: number
}

// noinspection JSUnusedGlobalSymbols
export interface ContentSettingsFilter {

    ContentType: string,
    ContentIds: Array<any>,
    ComponentSettingId: number,
    ParentId: number,
    ComponentSettingIds: Array<any>,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Country {

    Id: number,
    Name: string,
    Code: string,
    Languages: Array<any>,
    CokCode: string,
    IsImportCity: false,
    Localname: string,
    FlagImageId: number
}

// noinspection JSUnusedGlobalSymbols
export interface CountryFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface CountryLanguage {

    Id: number,
    Country: {},
    Language: {},
    CountryId: number,
    LanguageId: number,
    IsDefault: false
}

// noinspection JSUnusedGlobalSymbols
export interface CountryLanguageFilter {

    CountryIds: Array<any>,
    CountryId: number,
    CountryCode: string,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface CountryMentor {

    Id: number,
    UserId: number,
    CountryId: number,
    LandingCode: string,
    OraId: number,
    CountryCode: string,
    Contract: string
}

// noinspection JSUnusedGlobalSymbols
export interface CountryMentorFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Currency {

    Id: number,
    Name: string,
    Symbol: string,
    Logo: string,
    IsoCode: number,
    LocalName: string
}

// noinspection JSUnusedGlobalSymbols
export interface CurrencyFilter {

    Codes: Array<any>,
    IsoCodes: Array<any>,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface CustomBlock {

    Id: number,
    Html: string,
    ProductId: number,
    Position: number,
    IsActive: false,
    ProductContentId: number
}

// noinspection JSUnusedGlobalSymbols
export interface CustomBlockFilter {

    ProductId: number,
    IsActive: false,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface DefaultRegionLanguage {

    Id: number,
    RegionId: number,
    LanguageId: number
}

// noinspection JSUnusedGlobalSymbols
export interface DefaultRegionLanguageFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface DeliveryMethod {

    Id: number,
    Name: string,
    IsActive: false,
    Code: string,
    MinDeliveryPeriod: number,
    DeliveryCost: number,
    ImageId: number,
    Image: {},
    OraCode: string
}

// noinspection JSUnusedGlobalSymbols
export interface DeliveryMethodFilter {

    CartId: number,
    CityId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number,
    IsRead: false
}

// noinspection JSUnusedGlobalSymbols
export interface DeliveryMethodPrice {

    Id: number,
    DeliveryMethodId: number,
    Weight: number,
    CityId: number,
    Price: number
}

// noinspection JSUnusedGlobalSymbols
export interface DeliveryMethodPriceFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface DeliveryMethodRegion {

    Id: number,
    DeliveryMethodId: number,
    RegionId: number,
    IsActive: false,
    DeliveryMethodRegionLanguages: Array<any>
}

// noinspection JSUnusedGlobalSymbols
export interface DeliveryMethodRegionFilter {

    DeliveryMethodId: number,
    IsActive: false,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface DeliveryMethodRegionLanguage {

    Id: number,
    DeliveryMethodRegionId: number,
    LanguageId: number
}

// noinspection JSUnusedGlobalSymbols
export interface DeliveryMethodRegionLanguageFilter {

    DeliveryMethodRegionId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface DeliveryMethodRegionLanguagePaymentMethod {

    Id: number,
    DeliveryMethodRegionLanguageId: number,
    PaymentMethodId: number
}

// noinspection JSUnusedGlobalSymbols
export interface DeliveryMethodRegionLanguagePaymentMethodFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface DeliveryPoint {

    Id: number,
    Name: string
}

// noinspection JSUnusedGlobalSymbols
export interface DeliveryPointFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface DeliveryPrice {

    Id: number,
    DeliveryMethodId: number,
    From: string,
    To: string,
    Mass: number,
    Price: number
}

// noinspection JSUnusedGlobalSymbols
export interface DeliveryPriceFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface DiffsData {

    ImportLogId: number,
    EntityId: number,
    Type: string,
    Diffs: Array<any>,
    Entity: {}
}

// noinspection JSUnusedGlobalSymbols
export interface DiffsDataFilter {

    ImportLogId: number,
    ImportStatus: string,
    FieldName: string,
    EntityId: number,
    IsImported: false,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface DiscountType {

    Id: number,
    Name: string,
    Code: string
}

// noinspection JSUnusedGlobalSymbols
export interface DiscountTypeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface District {

    Id: number,
    Name: string,
    OraId: number,
    Cities: Array<any>,
    CountryId: number,
    OraParishId: number
}

// noinspection JSUnusedGlobalSymbols
export interface DistrictFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Document {

    Id: number,
    Name: string,
    BaseId: number,
    RegionId: number,
    LanguageId: number,
    Caption: string,
    DocumentTypeId: number,
    DocumentType: {},
    Description: string,
    ExpiresAt: null,
    IsActive: false,
    BaseDocument: {}
}

// noinspection JSUnusedGlobalSymbols
export interface DocumentFilter {

    BaseId: number,
    File: {},
    Header: {},
    DocumentTypeId: number,
    BaseIds: Array<any>,
    ProductCodes: Array<any>,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface DocumentGroup {

    Id: number,
    Name: string,
    ParentId: number
}

// noinspection JSUnusedGlobalSymbols
export interface DocumentGroupFilter {

    ParentId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface DocumentType {

    Id: number,
    Name: string,
    Code: string
}

// noinspection JSUnusedGlobalSymbols
export interface DocumentTypeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Domain {

    Id: number,
    Name: string
}

// noinspection JSUnusedGlobalSymbols
export interface DomainFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface DomainImport {

    Id: number,
    Message: string,
    IsSuccess: false
}

// noinspection JSUnusedGlobalSymbols
export interface DomainImportFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface DropCache {

    Id: number,
    Code: string
}

// noinspection JSUnusedGlobalSymbols
export interface DropCacheFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface EntityStat {

    Id: number,
    Name: string,
    EntityModel: {},
    Insert: number,
    Update: number,
    Delete: number,
    Select: number,
    Errors: number,
    Period: {},
    ApplicationId: number,
    Records: number,
    Weight: number,
    CreatedAt: null
}

// noinspection JSUnusedGlobalSymbols
export interface EntityStatFilter {

    Period: {},
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ExternalLink {

    Id: number,
    Url: string,
    ExternalLinkTypeId: number,
    RegionId: number,
    LanguageId: number,
    ExternalLinkType: {},
    CountryId: number,
    IsMain: false
}

// noinspection JSUnusedGlobalSymbols
export interface ExternalLinkFilter {

    ExternalLinkTypeId: number,
    CountryId: number,
    IsMain: false,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ExternalLinkType {

    Id: number,
    OraCode: string
}

// noinspection JSUnusedGlobalSymbols
export interface ExternalLinkTypeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ExternalUserService {

    Id: number,
    Name: string
}

// noinspection JSUnusedGlobalSymbols
export interface ExternalUserServiceFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Faq {

    Id: number,
    RegionId: number,
    LanguageId: number,
    BaseId: number,
    Answer: string,
    Question: string,
    IsActive: false,
    BaseFaq: {}
}

// noinspection JSUnusedGlobalSymbols
export interface FaqFilter {

    RegionId: number,
    LanguageId: number,
    BaseId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface FaqGroup {

    Id: number,
    Name: string,
    ParentId: number
}

// noinspection JSUnusedGlobalSymbols
export interface FaqGroupFilter {

    ParentId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface FaqGroupLocale {

    Id: number,
    RegionId: number,
    LanguageId: number,
    FaqGroupId: number,
    Name: string,
    IsActive: false
}

// noinspection JSUnusedGlobalSymbols
export interface FaqGroupLocaleFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Feature {

    Id: number,
    Title: string,
    Description: string,
    FeatureContentTypeId: number,
    FeatureContentId: number,
    ProductId: number,
    SubTitle: string,
    TextArea: string,
    Video: {},
    Image: {}
}

// noinspection JSUnusedGlobalSymbols
export interface FeatureFilter {

    ProductIds: Array<any>,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface FeatureContentType {

    Id: number,
    Name: string,
    Code: string
}

// noinspection JSUnusedGlobalSymbols
export interface FeatureContentTypeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Feedback {

    Id: number,
    Text: string,
    Answer: string,
    AnsweredAt: null,
    IsActive: false,
    Rating: number,
    UserContract: string,
    OraId: number,
    LanguageId: number,
    Title: string,
    AuthorName: string,
    SourceName: string,
    IsShowMainPage: false,
    Product: {},
    PostedAt: null,
    ProductCode: number,
    SocialNetwork: string
}

// noinspection JSUnusedGlobalSymbols
export interface FeedbackFilter {

    IsActive: false,
    IsShowMainPage: false,
    IsProductExist: false,
    IsUniqueText: false,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface GoshaFilterIds {

    Ids: Array<any>,
    ExceptIds: Array<any>
}

// noinspection JSUnusedGlobalSymbols
export interface ForceUpdate {

    Id: number,
    ApplicationId: number
}

// noinspection JSUnusedGlobalSymbols
export interface ForceUpdateFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface FrontendEventFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface GroupProduct {

    Id: number,
    Name: string,
    RegionId: number,
    LanguageId: number,
    IsActive: false,
    Description: string
}

// noinspection JSUnusedGlobalSymbols
export interface GroupProductFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface HardUser {

    Id: number,
    UserId: number
}

// noinspection JSUnusedGlobalSymbols
export interface HardUserFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface I18n {

    Id: number,
    Code: string,
    BaseKeyword: string,
    Translate: string,
    TranslateWithVars: string,
    RegionId: number,
    LanguageId: number,
    IsTranslated: false,
    BaseTranslateVars: Array<any>,
    TranslateVars: Array<any>,
    TranslateId: number,
    CreatedAt: null,
    UpdatedAt: null
}

// noinspection JSUnusedGlobalSymbols
export interface I18nFilter {

    RegionCode: string,
    LanguageCode: string,
    Code: string,
    BaseTranslateId: number,
    WithVars: false,
    Codes: Array<any>,
    IsOnlyNotTranslated: false,
    IsOnlyTranslated: false,
    WithBaseKeyword: false,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Image {

    Id: number,
    Url: string,
    OrigUrl: string,
    TinyUrl: string,
    SmallUrl: string,
    MediumUrl: string,
    LargeUrl: string,
    TypeId: number,
    Size: number,
    Products: Array<any>
}

// noinspection JSUnusedGlobalSymbols
export interface ImageFilter {

    File: {},
    Header: {},
    TypeId: number,
    IsActive: false,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ImageProduct {

    Id: number,
    ImageId: number,
    ProductId: number,
    IsActive: false,
    Position: number,
    Image: {}
}

// noinspection JSUnusedGlobalSymbols
export interface ImageProductFilter {

    ImageId: number,
    ProductId: number,
    ImageIds: Array<any>,
    ProductIds: Array<any>,
    IsActive: false,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ImageType {

    Id: number,
    Name: string,
    Code: string
}

// noinspection JSUnusedGlobalSymbols
export interface ImageTypeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ImportData {

    Id: number,
    Entity: string,
    EntityId: number,
    Action: string,
    EntityIds: Array<any>
}

// noinspection JSUnusedGlobalSymbols
export interface ImportDataFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ImportDiff {

    Id: number,
    Type: string,
    EntityId: number,
    ImportStatus: string,
    RegionId: number,
    LanguageId: number,
    ImportLogId: number,
    FieldName: string,
    OraValue: string,
    CurValue: string,
    IsImported: false
}

// noinspection JSUnusedGlobalSymbols
export interface ImportDiffFilter {

    EntityId: number,
    EntityIds: Array<any>,
    ImportStatus: string,
    FieldName: string,
    ImportLogId: number,
    Types: Array<any>,
    Type: string,
    IsImported: false,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ImportDiffRange {

    Status: string,
    IsImported: false,
    Count: number
}

// noinspection JSUnusedGlobalSymbols
export interface ImportDiffRangeFilter {

    EntityId: number,
    ImportStatus: string,
    RegionId: number,
    LanguageId: number,
    FieldName: string,
    ImportLogId: number,
    Type: string,
    IsImported: false,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ImportLog {

    Id: number,
    Type: string,
    Insert: number,
    Diff: number,
    Status: string,
    Error: string
}

// noinspection JSUnusedGlobalSymbols
export interface ImportLogFilter {

    Status: string,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Ingredient {

    Id: number,
    RegionId: number,
    LanguageId: number,
    BaseIngredientId: number,
    IsActive: false,
    ImageId: number,
    Name: string,
    Description: string,
    DescriptionShort: string,
    IngredientGroupId: number,
    Image: {}
}

// noinspection JSUnusedGlobalSymbols
export interface IngredientFilter {

    BaseIngredientId: number,
    IsActive: false,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface IngredientGroup {

    Id: number,
    Name: string,
    BaseIngredientGroupId: number,
    BaseIngredientGroup: {},
    RegionId: number,
    LanguageId: number
}

// noinspection JSUnusedGlobalSymbols
export interface IngredientGroupFilter {

    BaseIngredientGroupId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface LandingCompositionPercent {

    Id: number,
    Name: string,
    Description: string,
    RegionId: number,
    LanguageId: number,
    BaseLandingCompositionPercentId: number,
    BaseLandingCompositionPercent: {}
}

// noinspection JSUnusedGlobalSymbols
export interface LandingCompositionPercentFilter {

    BaseLandingCompositionPercentId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface LandingCompositionUnit {

    Id: number,
    Name: string,
    RegionId: number,
    LanguageId: number,
    BaseLandingCompositionUnitId: number,
    BaseLandingCompositionUnit: {}
}

// noinspection JSUnusedGlobalSymbols
export interface LandingCompositionUnitFilter {

    BaseLandingCompositionUnitId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Language {

    Id: number,
    Name: string,
    Code: string,
    NameRu: string,
    FlagImageId: number,
    FlagImage: {}
}

// noinspection JSUnusedGlobalSymbols
export interface LanguageFilter {

    Code: string,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Layout {

    Id: number,
    Name: string,
    RegionId: number,
    LanguageId: number,
    Uri: string,
    LayoutTypeId: number,
    LayoutType: {}
}

// noinspection JSUnusedGlobalSymbols
export interface LayoutFilter {

    RegionId: number,
    LanguageId: number,
    IsNotLoadSeoMeta: false,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface LayoutContent {

    Id: number,
    LayoutId: number,
    ComponentTemplateCode: string,
    Position: number,
    IsActive: false,
    RegionId: number,
    LanguageId: number,
    Name: string,
    LayoutTypeSlotId: number,
    LayoutTypeSlot: {}
}

// noinspection JSUnusedGlobalSymbols
export interface LayoutContentFilter {

    LayoutId: number,
    RegionId: number,
    LanguageId: number,
    IsActive: false,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface LayoutType {

    Id: number,
    Name: string,
    Code: string
}

// noinspection JSUnusedGlobalSymbols
export interface LayoutTypeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface LayoutTypeSlot {

    Id: number,
    Code: string,
    LayoutTypeId: number
}

// noinspection JSUnusedGlobalSymbols
export interface LayoutTypeSlotFilter {

    LayoutTypeId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface LiveShoppingGroup {

    Id: number,
    Name: string,
    RegionId: number,
    IsActive: false,
    LanguageId: number
}

// noinspection JSUnusedGlobalSymbols
export interface LiveShoppingGroupFilter {

    LanguageId: number,
    RegionId: number,
    IsActive: false,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface LiveShoppingGroupProduct {

    Id: number,
    LiveShoppingGroupId: number,
    ProductId: number,
    RegionId: number,
    IsActive: false,
    Region: {},
    Position: number,
    Product: {}
}

// noinspection JSUnusedGlobalSymbols
export interface LiveShoppingGroupProductFilter {

    LiveShoppingGroupId: number,
    ProductId: number,
    RegionId: number,
    IsActive: false,
    LiveShoppingGroupIds: Array<any>,
    ProductIds: Array<any>,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface MediaPublication {

    Id: number,
    Position: number,
    Title: string,
    Url: string,
    BaseMediaPublicationId: number,
    BaseMediaPublication: {},
    LanguageId: number,
    RegionId: number,
    ImageId: number,
    Image: {},
    Description: string
}

// noinspection JSUnusedGlobalSymbols
export interface MediaPublicationFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface MediaPublicationProduct {

    Id: number,
    MediaPublicationId: number,
    ProductId: number,
    MediaPublication: {}
}

// noinspection JSUnusedGlobalSymbols
export interface MediaPublicationProductFilter {

    ProductId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Mentor {

    Id: number,
    Contract: string,
    Email: string,
    FirstName: string,
    LastName: string,
    FamilyName: string,
    Icon: {}
}

// noinspection JSUnusedGlobalSymbols
export interface MentorFilter {

    Contract: string,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Menu {

    Id: number,
    Name: string,
    RegionId: number,
    LanguageId: number,
    MenuItems: Array<any>
}

// noinspection JSUnusedGlobalSymbols
export interface MenuFilter {

    RegionId: number,
    LanguageId: number,
    Name: string,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface MenuItem {

    Id: number,
    MenuId: number,
    Name: string,
    Url: string,
    IsActive: false,
    ColorPrimary: string,
    ColorSecondary: string,
    IsUndrelined: false,
    ParentId: number,
    Description: string,
    GroupNumber: number,
    Sort: number,
    SuperscriptText: string,
    SuperscriptColor: string,
    ImageId: number,
    Title: string,
    SeometaId: number,
    SeoMeta: {},
    Image: {},
    FrontendEvents: Array<any>,
    MenuItemTypeMobileId: number,
    MenuItemTypeDesktopId: number,
    MenuItemTypeMobile: {},
    MenuItemTypeDesktop: {},
    IsShowInBreadCrumb: false
}

// noinspection JSUnusedGlobalSymbols
export interface MenuItemFilter {

    MenuId: number,
    ParentId: number,
    GroupNumber: number,
    SeometaId: number,
    IsNotUpdateSort: false,
    MenuItemTypeMobileId: number,
    MenuItemTypeDesktopId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface MenuItemFrontendEvent {

    Id: number,
    MenuItemId: number,
    FrontendEventId: number
}

// noinspection JSUnusedGlobalSymbols
export interface MenuItemFrontendEventFilter {

    MenuItemId: number,
    FrontendEventId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface MenuItemType {

    Id: number,
    Name: string,
    Code: string
}

// noinspection JSUnusedGlobalSymbols
export interface MenuItemTypeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface MetroStation {

    Id: number,
    Name: string,
    Color: string,
    CityId: number
}

// noinspection JSUnusedGlobalSymbols
export interface MetroStationFilter {

    CityId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Migration {

    Id: number,
    MigrationId: number,
    Name: string,
    Hash: string
}

// noinspection JSUnusedGlobalSymbols
export interface MigrationFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface MiniProduct {

    Id: number,
    Code: number,
    RegionId: number,
    LanguageId: number,
    Model3dId: number,
    ImageIds: Array<any>,
    BaseDocumentIds: Array<any>,
    ImageProduct: {}
}

// noinspection JSUnusedGlobalSymbols
export interface MiniProductFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Model3d {

    Id: number,
    Code: string,
    Images: Array<any>,
    UpdatedAt: null,
    CreatedAt: null,
    IsActive: false,
    Products: Array<any>
}

// noinspection JSUnusedGlobalSymbols
export interface Model3dFilter {

    IsActive: false,
    ProductCodes: Array<any>,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Model3dImages {

    Id: number,
    Model3dId: number,
    ImageId: number
}

// noinspection JSUnusedGlobalSymbols
export interface Model3dImagesFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface MyCountry {

    Id: number,
    CountryCode: string,
    CountryId: number,
    RegionId: number,
    DefaultLanguage: {},
    DefaultCurrency: {}
}

// noinspection JSUnusedGlobalSymbols
export interface MyCountryFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface News {

    Id: number,
    RegionId: number,
    LanguageId: number,
    Text: string,
    Title: string,
    PreviewImageId: number,
    SeoTitle: string,
    SeoDescription: string,
    OraId: number,
    OraUrl: string,
    DateCreated: null,
    DateUpdated: null,
    BaseNewsId: number,
    IsPublished: false,
    IsMain: false,
    PreviewImage: {},
    Blog: {},
    Topic: {},
    BaseNewsOraId: number,
    RegionOraCode: string,
    PreviewImageUrl: string,
    LanguageOraCode: string,
    SeoMeta: {}
}

// noinspection JSUnusedGlobalSymbols
export interface NewsFilter {

    IsPreview: false,
    IsPublished: false,
    DateCreated: null,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface NewsletterSubscription {

    Id: number,
    Email: string,
    UserId: null
}

// noinspection JSUnusedGlobalSymbols
export interface NewsletterSubscriptionFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Oauth {

    Url: string,
    ServiceCode: string,
    Code: string,
    Auth: {}
}

// noinspection JSUnusedGlobalSymbols
export interface OauthFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number,
    Domain: string
}

// noinspection JSUnusedGlobalSymbols
export interface OauthUser {

    Id: number,
    ServiceCode: string,
    UserId: number,
    ServiceUserId: string,
    Code: string,
    User: {}
}

// noinspection JSUnusedGlobalSymbols
export interface OauthUserFilter {

    Domain: string,
    ServiceUserId: string,
    ServiceCode: string,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Opinion {

    Id: number,
    ProductId: number,
    Header: string,
    Quote: string,
    Author: string,
    LinkToSource: string,
    ImageId: number,
    Image: {},
    Position: number,
    RegionId: number,
    LanguageId: number,
    VideoId: number,
    Video: {}
}

// noinspection JSUnusedGlobalSymbols
export interface OpinionFilter {

    ProductId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface OraCountry {

    Name: string,
    Localname: string,
    CountryCode: string,
    Langs: string,
    Sclad: string
}

// noinspection JSUnusedGlobalSymbols
export interface OraCountryFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface OraImport {

    Name: string,
    Code: string,
    IsActive: false,
    Percent: number,
    StartAt: null
}

// noinspection JSUnusedGlobalSymbols
export interface OraImportFilter {

    IntervalName: string,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface OraLanguage {

    Name: string,
    NameRu: string,
    Code: string
}

// noinspection JSUnusedGlobalSymbols
export interface OraLanguageFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface OraProduct {

    Id: number,
    IsActive: false,
    Code: number,
    RegionId: number,
    LanguageId: number,
    IsBad: false,
    Composition: string,
    Contraindications: string,
    Description: string,
    DescriptionShort: string,
    GoogleCatId: number,
    OraId: number,
    HasAdvLicense: false,
    IsArchived: false,
    Brief: string,
    CEODescrypt: string,
    CEOHeader: string,
    CEOKeywords: string,
    Name: string,
    NameShort: string,
    NStatusId: null,
    IsOnlyOnline: false,
    ShowBrief: false,
    IsSiteShow: false,
    Sposob: string,
    UrlCode: string,
    Status: false,
    Price: number,
    CurrencyId: number,
    NStatus: {},
    ImportStatus: string,
    ImportLogId: number,
    IsShowDescription: false
}

// noinspection JSUnusedGlobalSymbols
export interface OraProductFilter {

    LanguageId: number,
    RegionId: number,
    ImportStatus: string,
    ImportLogId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Order {

    Id: number,
    UserId: number,
    TotalCost: number,
    TotalPoint: number,
    IsPackageMode: false,
    DeliveryMethodId: number,
    DeliveryCost: number,
    ProductsCost: number,
    AddressId: number,
    CartId: number,
    Bonuses: number,
    PromoCode: string,
    Credit: number,
    PaymentMethodId: number,
    OraCartId: number,
    OraHeaderId: number,
    IsPaid: false,
    RegionId: number,
    DomainName: string,
    ClientServiceCenterId: number,
    RegionCode: string,
    OraAddressId: number,
    OrderPackages: Array<any>,
    OrderTypeId: number,
    CityId: number,
    OrderDate: null,
    CurrencyId: number,
    IsAlternativeUserContact: false,
    RemainingCost: number,
    Email: string,
    Currency: {},
    Address: {},
    DeliveryMethod: {},
    ClientServiceCenter: {},
    OrderPaymentSystems: Array<any>,
    DeliveryMethodCode: string,
    CurrencyCode: string,
    PaymentInfo: {}
}

// noinspection JSUnusedGlobalSymbols
export interface OrderFilter {

    UserId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface OrderPackage {

    Id: number,
    OrderId: number,
    RecipientUserId: number,
    RecipientUserContract: string,
    OrderPackageProducts: Array<any>
}

// noinspection JSUnusedGlobalSymbols
export interface OrderPackageFilter {

    OrderId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface OrderPackageProduct {

    Id: number,
    ProductId: number,
    OrderPackageId: number,
    Price: number,
    Point: number,
    Quantity: number,
    Product: {}
}

// noinspection JSUnusedGlobalSymbols
export interface OrderPackageProductFilter {

    IsPreloadProduct: false,
    OrderPackageId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface OrderPayRetry {

    Id: number,
    OrderId: number,
    UserId: number,
    HeaderId: number,
    IsSuccess: false,
    ErrorMessage: string,
    PaymentInfo: {}
}

// noinspection JSUnusedGlobalSymbols
export interface OrderPayRetryFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface OrderPaymentSystem {

    Id: number,
    OrderId: number,
    PaymentSystemId: number,
    PaymentSystemTxId: string,
    Amount: number,
    Result: string,
    IsSuccess: false,
    PaymentAt: null,
    PaymentMethod: {}
}

// noinspection JSUnusedGlobalSymbols
export interface OrderPaymentSystemFilter {

    OrderId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface OrderProduct {

    Id: number,
    OrderId: number,
    ProductId: number,
    Count: number
}

// noinspection JSUnusedGlobalSymbols
export interface OrderProductFilter {

    OrderId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface OrderType {

    Id: number,
    Name: string
}

// noinspection JSUnusedGlobalSymbols
export interface OrderTypeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface PMarkers {

    Id: number,
    Name: string,
    Code: string
}

// noinspection JSUnusedGlobalSymbols
export interface PMarkersFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface PStatus {

    Id: number,
    Name: string
}

// noinspection JSUnusedGlobalSymbols
export interface PStatusFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface PageContent {

    Id: number,
    Name: string,
    PageTemplateId: number,
    IsActive: false,
    RegionId: number,
    LanguageId: number,
    Position: number,
    ComponentTemplateCode: string,
    LayoutTypeSlotId: number,
    LayoutTypeSlot: {}
}

// noinspection JSUnusedGlobalSymbols
export interface PageContentFilter {

    LanguageId: number,
    RegionId: number,
    PageId: number,
    IsActive: false,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface PageInfo {

    SeoMeta: {},
    LayoutContent: Array<any>,
    PageContent: Array<any>,
    Components: Array<any>,
    Settings: Array<any>,
    ProductContent: Array<any>,
    Layout: {}
}

// noinspection JSUnusedGlobalSymbols
export interface PageInfoFilter {

    Chpu: string,
    PageTemplateId: number,
    LayoutId: number,
    RegionId: number,
    LanguageId: number,
    Host: string,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface PageTemplate {

    Id: number,
    PageTypeId: number,
    Name: string,
    IsActive: false,
    RegionId: number,
    LanguageId: number,
    RootPageId: number,
    LayoutId: number,
    Uri: string,
    EntityId: number,
    IsDefault: false
}

// noinspection JSUnusedGlobalSymbols
export interface PageTemplateFilter {

    LanguageId: number,
    RegionId: number,
    PageTypeId: number,
    EntityId: number,
    IsNotLoadSeoMeta: false,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface PageType {

    Id: number,
    Name: string,
    Code: string
}

// noinspection JSUnusedGlobalSymbols
export interface PageTypeFilter {

    Code: string,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface PaymentCallback {

    Id: number,
    Response: string
}

// noinspection JSUnusedGlobalSymbols
export interface PaymentCallbackFilter {

    PaymentSystemName: string,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface PaymentMethod {

    Id: number,
    Name: string,
    IsActive: false,
    Code: string,
    ImageId: number,
    Image: {}
}

// noinspection JSUnusedGlobalSymbols
export interface PaymentMethodFilter {

    DeliveryMethodId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface PaymentMethodRegion {

    Id: number,
    PaymentMethodId: number,
    RegionId: number,
    IsActive: false,
    PaymentMethodRegionLanguages: Array<any>
}

// noinspection JSUnusedGlobalSymbols
export interface PaymentMethodRegionFilter {

    PaymentMethodId: number,
    RegionId: number,
    IsActive: false,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface PaymentMethodRegionLanguage {

    Id: number,
    PaymentMethodRegionId: number,
    LanguageId: number
}

// noinspection JSUnusedGlobalSymbols
export interface PaymentMethodRegionLanguageFilter {

    PaymentMethodRegionId: number,
    LanguageId: number,
    RegionId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface PortalAuth {

    Cookie: string
}

// noinspection JSUnusedGlobalSymbols
export interface PortalAuthFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Product {

    Id: number,
    IsActive: false,
    Code: number,
    RegionId: number,
    LanguageId: number,
    IsBad: false,
    Composition: string,
    Contraindications: string,
    Description: string,
    DescriptionShort: string,
    GoogleCatId: number,
    OraId: number,
    HasAdvLicense: false,
    IsArchived: false,
    Brief: string,
    CEODescrypt: string,
    CEOHeader: string,
    CEOKeywords: string,
    Name: string,
    NameShort: string,
    NStatusId: null,
    IsOnlyOnline: false,
    ShowBrief: false,
    IsSiteShow: false,
    UrlCode: string,
    Status: false,
    Markers: Array<any>,
    NStatus: {},
    Region: {},
    Language: {},
    Price: number,
    CurrencyId: number,
    Currency: {},
    Images: Array<any>,
    Videos: Array<any>,
    EnergyValue: number,
    Proteins: number,
    Fats: number,
    Carbohydrates: number,
    Packaged: string,
    Point: number,
    OldPrice: number,
    Reviews: Array<any>,
    Faqs: Array<any>,
    Documents: Array<any>,
    Advantages: Array<any>,
    ProductSaldo: {},
    ProductIcons: Array<any>,
    Ingredients: Array<any>,
    Strainers: Array<any>,
    Discount: number,
    DiscountPercentage: number,
    Chips: number,
    Weight: number,
    FullComposition: string,
    ImageLabel: {},
    ImageLabelId: number,
    Stars: number,
    ProductOptions: Array<any>,
    UseWayText: string,
    UseWayImageId: number,
    UseWayImage: {},
    UseWayVideoId: number,
    UseWayVideo: {},
    ChooseGroupId: number,
    SaleEndDate: null,
    SaleStartDate: null,
    FeedbackCount: number,
    GarantQualityText: string,
    UseWayVideoTitle: string,
    UseWayHeaderText: string,
    AnotherLanguageProduct: Array<any>,
    Model3dId: number,
    Model3d: {},
    ImportSourceRegionId: number,
    ImportSourceRegion: {},
    LandingTitle: string,
    LandingDescription: string,
    RecipeHeader: string,
    QualityAssuranceId: number,
    QualityAssurance: {},
    ConsumerTitle: string,
    ConsumerDescription: string,
    ProvenByScienceTitle: string,
    UseWayHackContent: string,
    UseWayVideoUrl: string,
    UseWayIsImageVideoPreview: false,
    LandingCompositionTitle: string,
    LandingCompositionPromoImageId: number,
    LandingCompositionPromoImage: {},
    LandingCompositionDescriptionText: string,
    LandingCompositionUseDescriptionText: false,
    ProductLandingSettingsShowOnSite: false,
    ProductLandingSettingsAccentColor: string,
    Features: Array<any>,
    ProductLandingCompositions: Array<any>,
    QualityMarks: Array<any>
}

// noinspection JSUnusedGlobalSymbols
export interface ProductByRegion {

    Id: number,
    IsActive: false,
    Code: number,
    RegionId: number,
    LanguageId: number,
    IsBad: false,
    Composition: string,
    Contraindications: string,
    Description: string,
    DescriptionShort: string,
    GoogleCatId: number,
    OraId: number,
    HasAdvLicense: false,
    IsArchived: false,
    Brief: string,
    CEODescrypt: string,
    CEOHeader: string,
    CEOKeywords: string,
    Name: string,
    NameShort: string,
    NStatusId: null,
    IsOnlyOnline: false,
    ShowBrief: false,
    IsSiteShow: false,
    Sposob: string,
    UrlCode: string,
    Status: false,
    Markers: Array<any>,
    NStatus: {},
    Region: {},
    Language: {},
    AnotherLanguageProduct: Array<any>
}

// noinspection JSUnusedGlobalSymbols
export interface ProductFilter {

    RegionId: number,
    LanguageId: number,
    Code: number,
    Codes: Array<any>,
    IsActive: false,
    IsSearch: false,
    SaldoCityIds: Array<any>,
    CityId: number,
    IsAvailable: false,
    Model3dIds: Array<any>,
    WithAnotherLanguages: false,
    ProductIconIds: Array<any>,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number,
    OraId: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductAddition {

    Id: number,
    AdditionProductId: number,
    ProductId: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductAdditionFilter {

    AdditionProductId: number,
    ProductId: number,
    AdditionProductIds: Array<any>,
    ProductIds: Array<any>,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductAdvantage {

    Id: number,
    ProductId: number,
    AdvantageId: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductAdvantageFilter {

    ProductId: number,
    AdvantageId: number,
    ProductIds: Array<any>,
    AdvantageIds: Array<any>,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductAffiliatedProduct {

    Id: number,
    ProductId: number,
    AffiliatedProductId: number,
    Position: number,
    AffiliatedProduct: {}
}

// noinspection JSUnusedGlobalSymbols
export interface ProductAffiliatedProductFilter {

    ProductId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductAlternative {

    Id: number,
    Name: string,
    RegionId: number,
    LanguageId: number,
    ProductAlternativeItems: Array<any>,
    ProductAlternativeTypeId: number,
    ProductAlternativeValueId: number,
    ProductAlternativeType: {},
    ProductAlternativeValue: {},
    ParentId: number,
    ImageId: number,
    Image: {},
    IconId: number,
    Icon: {},
    ProductAlternativeDisplayTypeId: number,
    ProductAlternativeDisplayType: {}
}

// noinspection JSUnusedGlobalSymbols
export interface ProductAlternativeFilter {

    ParentId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductAlternativeDisplayType {

    Id: number,
    Name: string,
    Code: string
}

// noinspection JSUnusedGlobalSymbols
export interface ProductAlternativeDisplayTypeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductAlternativeForProduct {

    Id: number,
    ProductAlternatives: Array<any>
}

// noinspection JSUnusedGlobalSymbols
export interface ProductAlternativeForProductFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductAlternativeItem {

    Id: number,
    ProductAlternativeId: number,
    ProductId: number,
    ProductAlternativeValueId: number,
    ProductAlternativeValue: {},
    ImageId: number,
    Image: {},
    Icon: {},
    IconId: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductAlternativeItemFilter {

    ProductId: number,
    ProductAlternativeId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductAlternativeType {

    Id: number,
    Name: string,
    Code: string
}

// noinspection JSUnusedGlobalSymbols
export interface ProductAlternativeTypeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductAlternativeValue {

    Id: number,
    ProductAlternativeTypeId: number,
    Value: string,
    Name: string
}

// noinspection JSUnusedGlobalSymbols
export interface ProductAlternativeValueFilter {

    ProductAlternativeTypeId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductCart {

    Id: number,
    ProductId: number,
    CartId: number,
    Quantity: number,
    Product: {},
    HiddenPrice: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductCartFilter {

    ProductId: number,
    CartId: number,
    CartIds: Array<any>,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductComplexProgram {

    Id: number,
    ProductId: number,
    ComplexProgramId: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductComplexProgramFilter {

    ProductId: number,
    ComplexProgramId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductComposition {

    Id: number,
    CompositionId: number,
    ProductId: number,
    Header: string,
    PromoImageId: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductCompositionFilter {

    CompositionId: number,
    ProductId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductContent {

    Id: number,
    Name: string,
    IsActive: false,
    RegionId: number,
    LanguageId: number,
    Position: number,
    ComponentTemplateCode: string,
    LayoutTypeSlotId: number,
    ProductId: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductContentFilter {

    PageId: number,
    IsActive: false,
    ProductId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductCopyLandingProperty {

    Id: number,
    SourceProductId: number,
    DestinationProductId: number,
    IsFeature: false,
    IsConsumer: false,
    IsQualityMark: false,
    IsLandingComposition: false,
    IsUseWay: false,
    IsRecipe: false,
    IsMediaPublication: false,
    IsOpinion: false,
    IsQualityAssurance: false,
    IsCustomBlock: false,
    IsProvenByScience: false,
    IsAffiliatedProduct: false
}

// noinspection JSUnusedGlobalSymbols
export interface ProductCopyLandingPropertyFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductDocument {

    Id: number,
    DocumentId: number,
    ProductId: number,
    Document: {}
}

// noinspection JSUnusedGlobalSymbols
export interface ProductDocumentFilter {

    DocumentId: number,
    ProductId: number,
    ProductIds: Array<any>,
    DocumentIds: Array<any>,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductFaq {

    Id: number,
    FaqId: number,
    ProductId: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductFaqFilter {

    ProductId: number,
    FaqId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductFavorite {

    Id: number,
    ProductId: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductFavoriteFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductFeedback {

    Id: number,
    ProductId: number,
    FeedbackId: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductFeedbackFilter {

    ProductId: number,
    FeedbackId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductGroupProduct {

    Id: number,
    ProductId: number,
    GroupId: number,
    Product: {}
}

// noinspection JSUnusedGlobalSymbols
export interface ProductGroupProductFilter {

    ProductId: number,
    GroupId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductIcon {

    Id: number,
    ProductId: number,
    BaseProductIconId: number,
    Name: string,
    DateStart: null,
    DateEnd: null,
    OraId: number,
    BaseProductIcon: {}
}

// noinspection JSUnusedGlobalSymbols
export interface ProductIconFilter {

    ProductId: number,
    BaseProductIconId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductImport {

    Id: number,
    UserId: number,
    StatusId: number,
    ProductImportStatusError: string,
    SourceOraCountryCode: string,
    DestinationRegionId: number,
    CopySuccessful: number,
    CopyErrors: number,
    CopyIgnored: number,
    IsIgnoreIfExists: false,
    ProductImportSourceAppId: number,
    ProductImportCodes: Array<any>,
    ProductImportLanguages: Array<any>,
    DestinationRegion: {},
    CreatedAt: null,
    UpdatedAt: null
}

// noinspection JSUnusedGlobalSymbols
export interface ProductImportFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductImportCode {

    Id: number,
    ProductCode: string,
    ProductImportId: number,
    ProductImportCodeStatusId: number,
    ProductImportCodeStatusError: string
}

// noinspection JSUnusedGlobalSymbols
export interface ProductImportCodeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductImportCodeStatus {

    Id: number,
    Name: string,
    Description: string
}

// noinspection JSUnusedGlobalSymbols
export interface ProductImportCodeStatusFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductImportDispatcher {

    Id: number,
    ProductImportId: number,
    ProductImportStatusId: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductImportDispatcherFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductImportLanguage {

    Id: number,
    ProductImportId: number,
    DestinationLanguageId: number,
    SourceLanguageCode: string,
    DestinationLanguage: {}
}

// noinspection JSUnusedGlobalSymbols
export interface ProductImportLanguageFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductImportSourceApp {

    Id: number,
    Name: string,
    Description: string
}

// noinspection JSUnusedGlobalSymbols
export interface ProductImportSourceAppFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductImportStatus {

    Id: number,
    Name: string,
    Description: string
}

// noinspection JSUnusedGlobalSymbols
export interface ProductImportStatusFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductIngredient {

    Id: number,
    IngredientId: number,
    ProductId: number,
    Ingredient: {}
}

// noinspection JSUnusedGlobalSymbols
export interface ProductIngredientFilter {

    IngredientId: number,
    ProductId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductLandingComposition {

    Id: number,
    ProductId: number,
    Position: number,
    Title: string,
    Description: string,
    LandingCompositionPercentId: number,
    LandingCompositionPercent: {},
    LandingCompositionUnitId: number,
    LandingCompositionUnit: {},
    Value: string
}

// noinspection JSUnusedGlobalSymbols
export interface ProductLandingCompositionFilter {

    ProductId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductLink {

    Id: number,
    MasterId: number,
    SlaveId: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductLinkFilter {

    MasterId: number,
    SlaveId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductMarker {

    Id: number,
    ProductId: number,
    MarkerId: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductMarkerFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductOption {

    Id: number,
    ProductId: number,
    Name: string,
    Value: string,
    OraOptionId: number,
    ProductOptionType: {},
    ProductOptionBaseTypeId: number,
    ProductOptionBaseValueId: number,
    IsSpecialName: false,
    IsSpecialValue: false,
    IsActive: false,
    IsSpecialShowInCatalog: false,
    IsSpecialShowInCatalogValue: false
}

// noinspection JSUnusedGlobalSymbols
export interface ProductOptionFilter {

    ProductId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductOptionBase {

    Id: number,
    ProductCode: string,
    ProductOptionBaseTypeId: number,
    ProductOptionBaseValueId: number,
    RegionId: number,
    ProductOptionBaseType: {},
    ProductOptionBaseValue: {}
}

// noinspection JSUnusedGlobalSymbols
export interface ProductOptionBaseFilter {

    ProductCodes: Array<any>,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductOptionBaseType {

    Id: number,
    Name: string,
    IsShowInCatalog: false,
    OraId: number,
    ProductOptionTypes: Array<any>
}

// noinspection JSUnusedGlobalSymbols
export interface ProductOptionBaseTypeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductOptionBaseValue {

    Id: number,
    ProductOptionBaseTypeId: number,
    Value: string,
    OraId: number,
    ProductOptionValues: Array<any>
}

// noinspection JSUnusedGlobalSymbols
export interface ProductOptionBaseValueFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductOptionType {

    Id: number,
    Name: string,
    OraId: number,
    ProductOptionBaseTypeId: number,
    RegionId: number,
    LanguageId: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductOptionTypeFilter {

    ProductOptionBaseTypeIds: Array<any>,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductOptionValue {

    Id: number,
    ProductOptionBaseValueId: number,
    Value: string,
    RegionId: number,
    LanguageId: number,
    OraId: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductOptionValueFilter {

    ProductOptionBaseValueIds: Array<any>,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductRecipe {

    Id: number,
    RecipeId: number,
    ProductId: number,
    Sort: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductRecipeFilter {

    RecipeId: number,
    ProductId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductReplacement {

    Id: number,
    ReplacementProductId: number,
    ProductId: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductReplacementFilter {

    ReplacementProductId: number,
    ProductId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductReview {

    Id: number,
    ReviewId: number,
    ProductId: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductReviewFilter {

    ReviewId: number,
    ProductId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductSaldo {

    Id: number,
    ProductId: number,
    ScladId: number,
    Volume: number,
    Contract: string,
    ProductCode: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductSaldoFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductSameLink {

    Id: number,
    ProductId1: number,
    ProductId2: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductSameLinkFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductStrainer {

    Id: number,
    ProductId: number,
    StrainerId: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductStrainerFilter {

    StrainerId: number,
    ProductId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProductVideo {

    Id: number,
    ProductId: number,
    VideoId: number,
    Video: {}
}

// noinspection JSUnusedGlobalSymbols
export interface ProductVideoFilter {

    ProductId: number,
    VideoId: number,
    VideoIds: Array<any>,
    ProductIds: Array<any>,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Project {

    Id: number,
    Name: string,
    Description: string
}

// noinspection JSUnusedGlobalSymbols
export interface ProjectFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProjectApplicationSettings {

    Id: number,
    ProjectSettingsId: number,
    ApplicationId: number,
    Key: string,
    Value: string,
    RegionId: number,
    LanguageId: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProjectApplicationSettingsFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProjectServer {

    Id: number,
    ProjectId: number,
    ApplicationId: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProjectServerFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProjectSettings {

    Id: number,
    ProjectId: number,
    Name: string,
    Key: string,
    Value: string,
    IsArray: false,
    SettingsTypeId: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProjectSettingsFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProvenByScience {

    Id: number,
    Title: string,
    Subtitle: string,
    Url: string,
    Position: number,
    ProductId: number
}

// noinspection JSUnusedGlobalSymbols
export interface ProvenByScienceFilter {

    ProductId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface QualityAssurance {

    Id: number,
    LanguageId: number,
    Name: string,
    Html: string,
    RegionId: number,
    BaseQualityAssuranceId: number,
    BaseQualityAssurance: {}
}

// noinspection JSUnusedGlobalSymbols
export interface QualityAssuranceFilter {

    BaseQualityAssuranceId: number,
    BaseQualityAssuranceIds: Array<any>,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface QualityMark {

    Id: number,
    ProductId: number,
    ImageId: number,
    Image: {}
}

// noinspection JSUnusedGlobalSymbols
export interface QualityMarkFilter {

    ProductId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Rank {

    Id: number,
    Name: string,
    OraCode: number,
    ShortName: string,
    ImageId: number
}

// noinspection JSUnusedGlobalSymbols
export interface RankFilter {

    ImageId: number,
    Code: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Recipe {

    Id: number,
    Title: string,
    Calories: number,
    Proteins: number,
    Fats: number,
    Carbohydrates: number,
    VideoId: number,
    ImageId: number,
    Description: string,
    CookingTime: string,
    IsActive: false,
    RegionId: number,
    LanguageId: number,
    PreviewImageId: number,
    ShortName: string,
    CaloriesPer1numbernumberGrams: number,
    Video: {},
    Image: {},
    PreviewImage: {},
    IsComplete: false
}

// noinspection JSUnusedGlobalSymbols
export interface RecipeFilter {

    IsActive: false,
    IsComplete: false,
    IsNotComplete: false,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface RecipeIngredient {

    Id: number,
    RecipeId: number,
    Name: string,
    Value: number,
    Unit: string,
    IsActive: false,
    Sort: number
}

// noinspection JSUnusedGlobalSymbols
export interface RecipeIngredientFilter {

    RecipeId: number,
    IsActive: false,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface RecipeStep {

    Id: number,
    RecipeId: number,
    ImageId: number,
    Caption: string,
    IsActive: false,
    Sort: number,
    Image: {}
}

// noinspection JSUnusedGlobalSymbols
export interface RecipeStepFilter {

    RecipeId: number,
    IsActive: false,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Region {

    Id: number,
    Name: string,
    Code: string,
    Type: number,
    EntityId: number,
    Countries: Array<any>,
    FreeDeliveryTotal: number,
    CurrencyId: number,
    IsFreeDelivery: false,
    IsGlobalTotal: false,
    ContactEmailAddress: string,
    ServerUrl: string,
    FlagUrl: string,
    DefaultCityId: number,
    BonusPaymentLimit: number,
    Currency: {},
    DefaultCity: {},
    RegistrationTypeId: number,
    ConsultantRegistrationCost: number,
    CityMentorSearchStrategy: {},
    CityMentorSearchStrategyId: number,
    DefaultMentorContract: string,
    SeoMetaCreateStrategyId: number,
    IsImportActions: false,
    IsImportProductPrices: false
}

// noinspection JSUnusedGlobalSymbols
export interface RegionFilter {

    Code: string,
    Codes: Array<any>,
    LanguageId: number,
    IsPreloadUrl: false,
    CityMentorSearchStrategyId: number,
    IsImportActions: false,
    RegionId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface RegionCountry {

    Id: number,
    RegionId: number,
    CountryId: number,
    Country: {}
}

// noinspection JSUnusedGlobalSymbols
export interface RegionCountryFilter {

    RegionIds: Array<any>,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface RegionLanguage {

    Id: number,
    RegionId: number,
    LanguageId: number,
    Language: {}
}

// noinspection JSUnusedGlobalSymbols
export interface RegionLanguageFilter {

    LanguageCodes: Array<any>,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface RegionSetting {

    Id: number,
    RegionId: number,
    SettingTypeId: number,
    IntValue: number,
    FloatValue: number,
    StringValue: string,
    BoolValue: false,
    SettingType: {}
}

// noinspection JSUnusedGlobalSymbols
export interface RegionSettingFilter {

    SettingTypeId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface RegionSettingType {

    Id: number,
    Name: string,
    DataType: string,
    Code: string,
    DefaultIntValue: number,
    DefaultFloatValue: number,
    DefaultStringValue: string,
    DefaultBoolValue: false
}

// noinspection JSUnusedGlobalSymbols
export interface SettingTypeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface RegionType {

    Id: number,
    Name: string
}

// noinspection JSUnusedGlobalSymbols
export interface RegionTypeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Regionality {

    RegionId: number,
    LanguageId: number
}

// noinspection JSUnusedGlobalSymbols
export interface RegistrationType {

    Id: number,
    Name: string,
    Code: string
}

// noinspection JSUnusedGlobalSymbols
export interface RegistrationTypeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Resource {

    Id: number,
    Name: string,
    Code: string,
    TypeId: number
}

// noinspection JSUnusedGlobalSymbols
export interface ResourceFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ResourceType {

    Id: number,
    Name: string
}

// noinspection JSUnusedGlobalSymbols
export interface ResourceTypeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface APIStatus {

    Status: string
}

// noinspection JSUnusedGlobalSymbols
export interface APIError {

    Error: false,
    Errors: Array<any>
}

// noinspection JSUnusedGlobalSymbols
export interface Error {

    ErrorMessage: string,
    ErrorCode: number,
    Field: string,
    ErrorDebug: string
}

// noinspection JSUnusedGlobalSymbols
export interface Pagination {

    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Review {

    Id: number,
    RegionId: number,
    LanguageId: number,
    BaseId: number,
    Text: string,
    Caption: string,
    ImgUrl: string,
    Title: string
}

// noinspection JSUnusedGlobalSymbols
export interface ReviewFilter {
    BaseId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface ReviewGroup {

    Id: number,
    Name: string,
    ParentId: number
}

// noinspection JSUnusedGlobalSymbols
export interface ReviewGroupFilter {

    ParentId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Role {

    Id: number,
    Name: string,
    Description: string
}

// noinspection JSUnusedGlobalSymbols
export interface RoleFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface RoleResource {

    Id: number,
    RoleId: number,
    ResourceId: number,
    Find: false,
    Read: false,
    Create: false,
    Update: false,
    Delete: false,
    FindOrCreate: false,
    MaxPerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface RoleResourceFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface SearchConfig {

    Id: number,
    IsActive: false,
    EditorId: number,
    UpdatedAt: null,
    RegionId: number,
    LanguageId: number,
    SearchResults: Array<any>,
    SearchFocusHints: Array<any>
}

// noinspection JSUnusedGlobalSymbols
export interface SearchConfigFilter {

    IsOnlyCurrent: false,
    IsActive: false,
    WithLinkedEntities: false,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface SearchFHMiniBanner {

    Id: number,
    SearchFocusHintId: number,
    ImageUrl: string,
    Name: string,
    LinkUrl: string,
    IsActive: false
}

// noinspection JSUnusedGlobalSymbols
export interface SearchFHMiniBannerFilter {

    SearchFocusHintId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface SearchFHPromoHint {

    Id: number,
    SearchFocusHintId: number,
    Name: string,
    Url: string,
    IsActive: false
}

// noinspection JSUnusedGlobalSymbols
export interface SearchFHPromoHintFilter {

    SearchFocusHintId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface SearchFHPromoProduct {

    Id: number,
    SearchFocusHintId: number,
    ProductId: number,
    IsActive: false
}

// noinspection JSUnusedGlobalSymbols
export interface SearchFHPromoProductFilter {

    SearchFocusHintId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface SearchFocusHint {

    Id: number,
    SearchConfigId: number,
    SearchFocusHintAreaId: number,
    Sort: number,
    IsActive: false,
    SearchFocusHintArea: {},
    SearchFHMiniBanners: Array<any>,
    SearchFHPromoHints: Array<any>,
    SearchFHPromoProducts: Array<any>
}

// noinspection JSUnusedGlobalSymbols
export interface SearchFocusHintFilter {

    SearchConfigId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface SearchFocusHintArea {

    Id: number,
    Name: string
}

// noinspection JSUnusedGlobalSymbols
export interface SearchFocusHintAreaFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface SearchQuery {

    Id: number,
    SearchQuery: string
}

// noinspection JSUnusedGlobalSymbols
export interface SearchQueryFilter {

    Query: string,
    StrainerIds: Array<any>,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface SearchQueryStrainer {

    Id: number,
    SearchQuery: string
}

// noinspection JSUnusedGlobalSymbols
export interface SearchQueryStrainerFilter {

    Query: string,
    StrainerIds: Array<any>,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface SearchResultArea {

    Id: number,
    Name: string
}

// noinspection JSUnusedGlobalSymbols
export interface SearchResultAreaFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Search {

    Id: number,
    Type: {},
    Title: string,
    Url: string,
    RegionId: number,
    LanguageId: number
}

// noinspection JSUnusedGlobalSymbols
export interface SearchResult {

    Id: number,
    SearchConfigId: number,
    IsActive: false,
    SearchResultAreaId: number,
    SearchResQnt: number,
    HintQnt: number,
    Sort: number
}

// noinspection JSUnusedGlobalSymbols
export interface SearchResultFilter {

    SearchConfigId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface SendContactEmail {

    Id: number,
    From: string,
    Text: string,
    HiddenField: string,
    RegionId: number,
    Name: string
}

// noinspection JSUnusedGlobalSymbols
export interface SendContactEmailFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface SeoMeta {

    Id: number,
    Title: string,
    Keywords: string,
    Description: string,
    RegionId: number,
    LanguageId: number,
    Chpu: string,
    PageTemplateId: null,
    EntityName: string,
    EntityId: null,
    RedirectSeoMetaId: number,
    RedirectSeoMeta: {},
    IsPermanentRedirect: false,
    PageTemplate: {}
}

// noinspection JSUnusedGlobalSymbols
export interface SeoMetaFilter {
    Chpu: string,
    EntityNames: Array<any>,
    EntityName: string,
    EntityIds: Array<any>,
    EntityId: number,
    PageTemplateId: number,
    PageTemplateIds: Array<any>,
    IsNotArchived: false,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface SeoMetaCreateStrategy {

    Id: number,
    Name: string,
    Description: string
}

// noinspection JSUnusedGlobalSymbols
export interface SeoMetaCreateStrategyFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Settings {

    Id: number,
    Name: string,
    Value: string,
    Description: string,
    GroupId: number,
    SettingsGroup: {}
}

// noinspection JSUnusedGlobalSymbols
export interface SettingsFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number,
    Name: string
}

// noinspection JSUnusedGlobalSymbols
export interface SettingsGroup {

    Id: number,
    Name: string
}

// noinspection JSUnusedGlobalSymbols
export interface SettingsGroupFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface SettingsType {

    Id: number,
    Code: string
}

// noinspection JSUnusedGlobalSymbols
export interface SettingsTypeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface StatPeriod {

    Id: number,
    Name: string,
    Value: {}
}

// noinspection JSUnusedGlobalSymbols
export interface StatPeriodFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Store {

    Id: number,
    Name: string,
    Address: string,
    Contract: string,
    Latitude: number,
    Longitude: number,
    CountryCode: string,
    CountryId: number,
    OraId: number
}

// noinspection JSUnusedGlobalSymbols
export interface StoreFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface StoreDeliveryMethod {

    Id: number,
    StoreId: number,
    DeliveryMethodId: number
}

// noinspection JSUnusedGlobalSymbols
export interface StoreDeliveryMethodFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Strainer {

    Id: number,
    StrainerTypeId: number,
    Value: string,
    StrainerType: {},
    BaseStrainerId: number,
    BaseStrainer: {}
}

// noinspection JSUnusedGlobalSymbols
export interface StrainerFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface StrainerType {

    Id: number,
    LanguageId: number,
    RegionId: number,
    BaseStrainerTypeId: number,
    Name: string,
    Strainers: Array<any>,
    BaseStrainerType: {}
}

// noinspection JSUnusedGlobalSymbols
export interface StrainerTypeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface StreetType {

    Id: number,
    Name: string,
    RegionId: number,
    LanguageId: number,
    BaseId: number,
    OraTypeId: number,
    RegionCode: string,
    LanguageCode: string
}

// noinspection JSUnusedGlobalSymbols
export interface StreetTypeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface SubscriptionProductStock {

    Id: number,
    Email: string,
    ProductId: number,
    IsSend: false,
    StoreId: number,
    AuthId: number
}

// noinspection JSUnusedGlobalSymbols
export interface SubscriptionProductStockFilter {

    AuthId: number,
    ProductId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface SystemTranslate {

    Id: number,
    RegionId: number,
    LanguageId: number,
    BaseSystemTranslateId: number,
    Text: string
}

// noinspection JSUnusedGlobalSymbols
export interface SystemTranslateFilter {

    BaseSystemTranslateId: number,
    BaseSystemTranslateIds: Array<any>,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Task {

    Id: number,
    Name: string,
    Code: string,
    TaskPeriodId: number,
    TaskPeriod: {}
}

// noinspection JSUnusedGlobalSymbols
export interface TaskFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface TaskPeriod {

    Id: number,
    Name: string,
    Interval: {}
}

// noinspection JSUnusedGlobalSymbols
export interface TaskPeriodFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Timezone {

    Id: number,
    Name: string,
    Abbr: string,
    Offset: number,
    IsDst: false,
    Text: string,
    Value: string
}

// noinspection JSUnusedGlobalSymbols
export interface TimezoneFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Topic {

    Id: number,
    Name: string,
    RegionId: number,
    LanguageId: number,
    BaseTopicId: number
}

// noinspection JSUnusedGlobalSymbols
export interface TopicFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Translate {

    Id: number,
    RegionId: number,
    LanguageId: number,
    Translate: string,
    Keyword: string,
    OraId: number,
    BaseTranslateId: number,
    Language: {},
    Region: {},
    BaseTranslate: {}
}

// noinspection JSUnusedGlobalSymbols
export interface TranslateFilter {

    BaseTranslateId: number,
    BaseTranslateIds: Array<any>,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface TranslateError {

    Id: number,
    Code: number,
    LanguageId: number,
    Translate: string
}

// noinspection JSUnusedGlobalSymbols
export interface TranslateErrorFilter {

    ErrorCodes: Array<any>,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface TranslateExport {

    Id: number,
    Comment: string
}

// noinspection JSUnusedGlobalSymbols
export interface TranslateExportFilter {

    Code: string,
    Codes: Array<any>,
    BaseTranslateId: number,
    RegionCode: string,
    LanguageCode: string,
    IsOnlyNotTranslated: false,
    IsOnlyTranslated: false,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface TranslateImport {

    Id: number,
    Comment: string
}

// noinspection JSUnusedGlobalSymbols
export interface TranslateImportFilter {

    File: {},
    Header: {},
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface TranslateList {

    Id: number,
    Keyword: string,
    Translate: string,
    List: Array<any>
}

// noinspection JSUnusedGlobalSymbols
export interface TranslatesData {

    Id: number,
    LanguageCode: string
}

// noinspection JSUnusedGlobalSymbols
export interface TranslateListFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface TranslateVar {

    Id: number,
    TranslateId: number,
    VarName: string,
    VarValue: string
}

// noinspection JSUnusedGlobalSymbols
export interface TranslateVarFilter {

    TranslateId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface UseWayDocument {

    Id: number,
    DocumentId: number,
    ProductId: number,
    Document: {}
}

// noinspection JSUnusedGlobalSymbols
export interface UseWayDocumentFilter {

    ProductId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface User {

    Id: number,
    Email: string,
    FirstName: string,
    IsActive: false,
    IsAuthorized: false,
    IsOnline: false,
    LastName: string,
    MobilePhone: string,
    Password: string,
    Contract: string,
    IconUrl: string,
    IsImported: false,
    ExternalUserServiceId: number,
    ExternalCode: string,
    CreatedAt: null,
    UpdatedAt: null,
    CashbackPercent: number,
    OauthUsers: Array<any>,
    CashBack: {},
    CashBacks: Array<any>,
    Point: number,
    Snowflakes: number,
    CreditLimit: number,
    CreditAvailable: false,
    CityId: number,
    MentorId: number,
    IsSubscribeNewsletter: false,
    RegistrationDomain: string,
    UserRoles: Array<any>,
    IconId: number,
    Icon: {},
    Gender: string,
    Birthday: null,
    FamilyName: string,
    Bonus: number,
    CurrencyId: number,
    Contacts: Array<any>,
    OraUserTypeId: number,
    OraRegistrationIsPaid: false,
    Roles: Array<any>,
    RankId: number
}

// noinspection JSUnusedGlobalSymbols
export interface UserShort {

    FirstName: string,
    LastName: string,
    Gender: string,
    FamilyName: string,
    Icon: {},
    Contacts: Array<any>
}

// noinspection JSUnusedGlobalSymbols
export interface UserFilter {

    Email: string,
    WithAuthorized: false,
    WithOnline: false,
    IsOnlyOnline: false,
    Contract: string,
    Contracts: Array<any>,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface UserAchievement {

    Id: number,
    UserId: number,
    AchievementId: number,
    UserContract: string,
    AchievementCode: string
}

// noinspection JSUnusedGlobalSymbols
export interface UserAchievementFilter {

    UserId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface UserActivate {

    Id: number,
    PhoneCode: string,
    EmailCode: string,
    Contract: string,
    Auth: {}
}

// noinspection JSUnusedGlobalSymbols
export interface UserActivateFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface UserArticle {

    Id: number,
    Text: string,
    Title: string,
    SubTitle: string,
    UserId: number,
    UserName: string,
    UserPhoto: {},
    CountryFlagImage: {},
    RankCode: number,
    UserAchievementCodes: Array<any>,
    UserContacts: Array<any>,
    DatePublished: null,
    VideoUrl: string,
    CountryName: string
}

// noinspection JSUnusedGlobalSymbols
export interface UserArticleContact {

    UserContactTypeId: number,
    Value: string
}

// noinspection JSUnusedGlobalSymbols
export interface UserArticleFilter {

    UserId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface UserCart {

    Id: number,
    Code: string,
    AuthId: number
}

// noinspection JSUnusedGlobalSymbols
export interface UserCartFilter {

    UserId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface UserContact {

    Id: number,
    UserId: number,
    UserContactTypeId: number,
    Value: string,
    UserContactType: {}
}

// noinspection JSUnusedGlobalSymbols
export interface UserContactFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface UserContactType {

    Id: number,
    Name: string
}

// noinspection JSUnusedGlobalSymbols
export interface UserContactTypeFilter {

    Name: string,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface UserMainContact {

    Id: number,
    Phone: string,
    Email: string,
    Contract: string,
    UserId: number,
    PhoneVerified: false,
    EmailVerified: false
}

// noinspection JSUnusedGlobalSymbols
export interface UserMainContactFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface UserMenu {

    Id: number,
    Name: string,
    RegionId: number,
    LanguageId: number,
    UserMenuItems: Array<any>
}

// noinspection JSUnusedGlobalSymbols
export interface UserMenuFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface UserMenuItem {

    Id: number,
    Name: string,
    UserMenuItemTypeId: number,
    UserMenuId: number,
    Url: string,
    Data: string,
    GroupNumber: number,
    ColorPrimary: string,
    ColorSecondary: string,
    Sort: number,
    Description: string,
    SeometaId: number,
    SeoMeta: {},
    FrontendEvents: Array<any>,
    UserMenuItemType: {},
    IsActive: false
}

// noinspection JSUnusedGlobalSymbols
export interface UserMenuItemFilter {

    UserMenuItemTypeId: number,
    UserMenuId: number,
    SeometaId: number,
    IsActive: false,
    IsNotUpdateSort: false,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface UserMenuItemFrontendEvent {

    Id: number,
    UserMenuItemId: number,
    FrontendEventId: number
}

// noinspection JSUnusedGlobalSymbols
export interface UserMenuItemFrontendEventFilter {

    UserMenuItemId: number,
    FrontendEventId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface UserMenuItemRole {

    Id: number,
    UserMenuItemId: number,
    RoleId: number
}

// noinspection JSUnusedGlobalSymbols
export interface UserMenuItemRoleFilter {

    UserMenuItemId: number,
    RoleId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface UserMenuItemType {

    Id: number,
    Name: string,
    Code: string
}

// noinspection JSUnusedGlobalSymbols
export interface UserMenuItemTypeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface UserPassword {

    Id: number,
    Code: string,
    Password: string,
    Contract: string,
    Email: string,
    Phone: string,
    Auth: {}
}

// noinspection JSUnusedGlobalSymbols
export interface UserPasswordFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface UserPasswordChange {

    Id: number,
    UserId: number,
    CurrentPassword: string,
    NewPassword: string
}

// noinspection JSUnusedGlobalSymbols
export interface UserPasswordChangeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface UserPasswordRestore {

    Id: number,
    UserId: number,
    Code: string,
    IsEmail: false,
    IsSms: false,
    IsActive: false,
    UsedDate: null,
    IsUsed: false,
    CreatedAt: null,
    DomainName: string,
    RestoreField: string
}

// noinspection JSUnusedGlobalSymbols
export interface UserPasswordRestoreFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface UserPasswordVerify {

    Id: number,
    Code: string
}

// noinspection JSUnusedGlobalSymbols
export interface UserPasswordVerifyFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface UserRegistration {

    Id: number,
    Name: string,
    MentorUserId: number,
    Email: string,
    Phone: string,
    Birthday: null,
    CityId: number,
    Gender: string,
    ClientServiceCenterId: number,
    IsSubscribeNewsletter: false,
    LandingCode: string,
    IsPrivileged: false,
    AuthId: number,
    Auth: {},
    IsFastReg: false,
    RegistrationDomain: string
}

// noinspection JSUnusedGlobalSymbols
export interface UserRegistrationFilter {

    IsNotIgnoreError: false,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface UserRegistrationOrder {

    Id: number,
    PaymentInfo: {},
    PaymentMethodId: number,
    DomainName: string,
    RegionId: number
}

// noinspection JSUnusedGlobalSymbols
export interface UserRegistrationOrderFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface UserRole {

    Id: number,
    RoleId: number,
    UserId: number,
    IsPaid: false
}

// noinspection JSUnusedGlobalSymbols
export interface UserRoleFilter {

    RoleId: number,
    UserId: number,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface UserStat {

    Id: number,
    WsToken: string,
    WsRequests: number,
    HttpRequests: number,
    CreateReq: number,
    ReadReq: number,
    UpdateReq: number,
    DeleteReq: number,
    FindOrCreateReq: number,
    Period: {},
    WsErrors: number,
    HttpErrors: number,
    UserId: number,
    ApplicationId: number,
    CreatedAt: null
}

// noinspection JSUnusedGlobalSymbols
export interface UserStatFilter {

    Period: {},
    DateStart: null,
    DateEnd: null,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface UserSynchronization {

    Id: number,
    UserId: number,
    AddressUpdatedAt: null,
    CertificateUpdatedAd: null,
    UserInfoUpdatedAt: null,
    OrderUpdatedAt: null
}

// noinspection JSUnusedGlobalSymbols
export interface UserSynchronizationFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface Video {

    Id: number,
    Url: string,
    OrigUrl: string,
    TinyUrl: string,
    SmallUrl: string,
    MediumUrl: string,
    LargeUrl: string,
    Name: string,
    VideoTypeId: number,
    OraId: string,
    IsActive: false
}

// noinspection JSUnusedGlobalSymbols
export interface VideoFilter {

    File: {},
    Header: {},
    FileName: string,
    IsActive: false,
    ProductCodes: Array<any>,
    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface VideoType {

    Id: number,
    Name: string,
    Code: string
}

// noinspection JSUnusedGlobalSymbols
export interface VideoTypeFilter {

    RegionId: number,
    LanguageId: number,
    Search: string,
    SearchBy: Array<any>,
    Order: Array<any>,
    OrderDirection: Array<any>,
    Ids: Array<any>,
    ExceptIds: Array<any>,
    CurrentPage: number,
    PerPage: number
}

// noinspection JSUnusedGlobalSymbols
export interface WssRequest {

    Route: string,
    Method: string,
    SessionId: string,
    Type: string,
    Data: {}
}
