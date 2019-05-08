const config = {};
const apiVersion = "v1";
export function setup({ url }) {
  config.url = url;
}

export async function list() {
  try {
    const response = await fetch(`${config.url}/${apiVersion}/todos`);
    if (response.status !== 200) {
      throw new Error(
        "Looks like there was a problem. Status Code: " + response.status
      );
    }
    const data = await response.json();
    return data;
  } catch (err) {
    throw err;
  }
}
export async function add({ title }) {
  try {
    const response = await fetch(`${config.url}/${apiVersion}/todos`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({ title })
    });
    if (response.status !== 201) {
      throw new Error(
        "Looks like there was a problem. Status Code: " + response.status
      );
    }
    const data = await response.json();
    return data;
  } catch (err) {
    throw err;
  }
}
export async function remove(id) {
  try {
    const response = await fetch(`${config.url}/${apiVersion}/todos/${id}`, {
      method: "DELETE"
    });
    if (response.status !== 204) {
      throw new Error(
        "Looks like there was a problem. Status Code: " + response.status
      );
    }
    return;
  } catch (err) {
    throw err;
  }
}
