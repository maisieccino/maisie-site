import GhostContentAPI from "@tryghost/content-api";
import { sys } from "typescript";
const ghostClient = new GhostContentAPI({
  url: "https://bell-blog.ghost.io",
  key: process.env.CONTENT_API_KEY || "",
  version: "v6.0",
});

if (posts == undefined) {
  console.error("no posts");
  sys.exit(1);
}
