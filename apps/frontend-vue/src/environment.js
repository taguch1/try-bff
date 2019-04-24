export default class Environment {
  static async setup() {
    const response = await fetch(
      `${location.protocol}//${location.host}/env`
    ).catch(err => {
      // throw err;
      throw new Error("failed to get env. err:" + err);
    });
    if (response.status !== 200) {
      throw new Error("failed to get env. Status Code: " + response.status);
    }
    const data = await response.json().catch(err => {
      throw new Error("failed to parse env. err:" + err);
    });
    Environment.config = data;
  }
  static get URL() {
    return Environment.config.url;
  }
}
