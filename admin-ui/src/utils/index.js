import configs from "../configs"

const apiUrl = {
  getProjectsUrl: () => `${configs.API_ENDPOINT}/projects`,
  postProjectUrl: () => `${configs.API_ENDPOINT}/projects`,
  getProjectUrl:  id => `${configs.API_ENDPOINT}/projects/${id}`,
  getFeaturesUrl: id => `${configs.API_ENDPOINT}/projects/${id}/features`,
  postFeatureUrl: id => `${configs.API_ENDPOINT}/projects/${id}/features`,
  getFeatureUrl: (id, key) => `${configs.API_ENDPOINT}/projects/${id}/features/${key}`,
  getUsersUrl: id => `${configs.API_ENDPOINT}/projects/${id}/users`,
  getUserUrl: (id, uuid) => `${configs.API_ENDPOINT}/projects/${id}/users/${uuid}`,
}

function fetchData(url) {
  return new Promise(async (resolve, reject) => {
    const res = await fetch(url)
    const data = await res.json()

    if (res.ok) resolve(data)

    reject(data.message)
  })
}

function postData(url, params) {
  return new Promise(async (resolve, reject) => {
    const options = {
      method: 'POST',
      mode: 'cors',
      body: JSON.stringify(params),
    }

    const res = await fetch(url, options)
    const data = await res.json()

    if (res.ok) resolve(data)

    reject(data.message)
  })
}


function formatDate(date) {
  date = typeof date === "object" ? d : new Date(d)
  const isValidDate = date.getFullYear() < new Date(0)
  if (!isValidDate) return "-"

  const y = date.getFullYear()
  const m = `0${date.getMonth() + 1}`.slice(-2)
  const d = `0${date.getDate()}`.slice(-2)

  return `${y}/${m}/${d}`
}

export {
  apiUrl,
  fetchData,
  postData,
  formatDate,
}
