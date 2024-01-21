var apm = require('elastic-apm-node').start({
  serviceName: 'node-app-1',

  secretToken: '<anything-you-want>', // from Step 6
  serverUrl: 'https://<fleet-server-domain>:8200',
  verifyServerCert: true,
  environment: 'production'
});

const express = require('express');
const app = express()
const port = 3030

app.get('/', (req, res) => {
  res.send('Hello World!')
})
app.get('/demo', (req, res) => {
  res.send('Welcome Demo')
})
app.get('/error', (req, res) => {
  res.send('Capture Error!')
  const err = new Error('Trigger Error!')
  apm.captureError(err)
})
app.listen(port, () => {
  console.log(`Listening on port ${port}`)
})
