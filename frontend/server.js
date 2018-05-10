const express = require('express')
const app = express()
const request = require('request');
var DateDiff = require('date-diff');

app.use(express.static('public'));
app.set('view engine', 'ejs')

let api = `http://localhost:1323`
let getAccidents = api + `/accidents`

app.get('/', function (req, res) {
    request(getAccidents, function (err, response, body) {
        if (err) {
            res.render('index', { data: null, error: 'Error, please try again' });
        } else {
            let accidents = JSON.parse(body)
            let latestAccidents = accidents[0]
            let diff = Date.diff(new Date(), new Date(accidents[0].date)).days()
            res.render('index', { accident: accidents[0], datediff: diff, error: null });
        }
    });
})

app.listen(3000, function () {
    console.log('Example app listening on port 3000!')
})