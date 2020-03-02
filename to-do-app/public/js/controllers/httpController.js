export class HTTP {
  //Make http GET request
  async get(url) {
    const response = await fetch(url);
    const resData = await response.json();
    return resData;
  }
}
