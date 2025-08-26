import * as wmill from "windmill-cli@1.481.0"

export async function main() {
  await wmill.hubPull({ workspace: "admins", token: process.env["WM_TOKEN"], baseUrl: globalThis.process.env["BASE_URL"] });
}
