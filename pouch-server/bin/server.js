const PouchDB = require("pouchdb");
const config = require("config");
const app = require("express-pouchdb")(PouchDB);
app.listen(config.get("port"));
