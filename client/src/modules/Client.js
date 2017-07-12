/* eslint-disable no-undef */
function search(query, cb) {
  return fetch(`api/tickets/${query}`, {
    accept: 'application/json',
  }).then(checkStatus)
    .then(parseJSON)
    .then(cb);
}

function listIndex(cb) {
  return fetch(`api/tickets`, {
    accept: 'application/json',
  }).then(checkStatus)
    .then(parseJSON)
    .then(cb);
}

function checkStatus(response) {
  if (response.status >= 200 && response.status < 300) {
    return response;
  } else {
    const error = new Error(`HTTP Error ${response.statusText}`);
    error.status = response.statusText;
    error.response = response;
    console.log(error); // eslint-disable-line no-console
    throw error;
  }
}

function parseHTML(response) {
  return JSON.stringify(response)
}

function parseJSON(response) {
  return response.json();
}

const Client = { search, listIndex };
export default Client;
