package cmd

const FindUrl = "/api/v1/{entity}"
const ReadUrl = "/api/v1/{entity}/" // + id
const CreateUrl = "/api/v1/{entity}"
const MultiCreateUrl = "/api/v1/{entity}/list"
const UpdateUrl = "/api/v1/{entity}/" // + id
const MultiUpdateUrl = "/api/v1/{entity}/list"
const DeleteUrl = "/api/v1/{entity}/" // + id
const MultiDeleteUrl = "/api/v1/{entity}/list"
const FindOrCreateUrl = "/api/v1/{entity}"
const UpdateOrCreateUrl = "/api/v1/{entity}"

const usualEntityJsRoute = `
export default {
    find: "` + FindUrl + `",
    read: "` + ReadUrl + `", // + id
    create: "` + CreateUrl + `",
    multiCreate: "` + MultiCreateUrl + `",
    update: "` + UpdateUrl + `", // + id
    multiUpdate: "` + MultiUpdateUrl + `",
    delete: "` + DeleteUrl + `", // + id
    multiDelete: "` + MultiDeleteUrl + `",
    findOrCreate: "` + FindOrCreateUrl + `",
    updateOrCreate: "` + UpdateOrCreateUrl + `"
};

`
