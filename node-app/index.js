const express = require('express');
const app = express();

app.get('/', (req, res) => {
  res.send('Hello from Node.js app!');
});

const PORT = 3000;
app.listen(PORT, () => {
  console.log(`Node.js server running on port ${PORT}`);
});
