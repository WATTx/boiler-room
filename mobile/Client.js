const endpoint = 'http://10.32.2.86:8080'

class Client {

  setLevel = (livelinkHost, lightId, level) => {
    const url = `${endpoint}/${livelinkHost}/lights/${lightId}`

    const opts = {
      method: 'PATCH',
      body: JSON.stringify({level: level}),
    }

    fetch(url, opts).then( res => {
      if (!res.ok) {
        return Promise.reject(new Error(`${res.status}: Can't set the light level`))
      }

    })
  }
}

export default Client
