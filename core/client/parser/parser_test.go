package parser

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/AleckDarcy/reload/core/tracer"

	"golang.org/x/net/html"
)

var str = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, shrink-to-fit=no">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Hipster Shop</title>
    <link href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.1/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-WskhaSGFgHYWDcbwN70/dfYBj47jz9qbsMId/iRN3ewGhXQFZCSftd1LZCfmhktB" crossorigin="anonymous">
</head>
<body>

    <header>
        <div class="navbar navbar-dark bg-dark box-shadow">
            <div class="container d-flex justify-content-between">
                <a href="/" class="navbar-brand d-flex align-items-center">
                    Hipster Shop
                </a>
                
                <form class="form-inline ml-auto" method="POST" action="/setCurrency" id="currency_form">
                    <select name="currency_code" class="form-control"
                    onchange="document.getElementById('currency_form').submit();" style="width:auto;">
                    
                        <option value="EUR" >EUR</option>
                    
                        <option value="USD" selected="selected">USD</option>
                    
                        <option value="JPY" >JPY</option>
                    
                        <option value="GBP" >GBP</option>
                    
                        <option value="TRY" >TRY</option>
                    
                        <option value="CAD" >CAD</option>
                    
                    </select>
                    <a class="btn btn-primary btn-light ml-2" href="/cart" role="button">View Cart (0)</a>
                </form>
                
            </div>
        </div>
    </header>



    <main role="main">
        <section class="jumbotron text-center mb-0"
		 
		>
            <div class="container">
                <h1 class="jumbotron-heading">
                    One-stop for Hipster Fashion &amp; Style Online
                </h1>
                <p class="lead text-muted">
                    Tired of mainstream fashion ideas, popular trends and
                    societal norms? This line of lifestyle products will help
                    you catch up with the hipster trend and express your
                    personal style. Start shopping hip and vintage items now!
                </p>
            </div>
        </section>

        <div class="py-5 bg-light">
            <div class="container">
            <div class="row">
                
                <div class="col-md-4">
                    <div class="card mb-4 box-shadow">
                        <a href="/product/OLJCESPC7Z">
                            <img class="card-img-top" alt =""
                                style="width: 100%; height: auto;"
                                src="/static/img/products/typewriter.jpg">
                        </a>
                        <div class="card-body">
                            <h5 class="card-title">
                                Vintage Typewriter
                            </h5>
                            <div class="d-flex justify-content-between align-items-center">
                                <div class="btn-group">
                                    <a href="/product/OLJCESPC7Z">
                                        <button type="button" class="btn btn-sm btn-outline-secondary">Buy</button>
                                    </a>
                                </div>
                                <small class="text-muted">
                                    USD 67.99 
                                </strong>
                                </small>
                            </div>
                        </div>
                    </div>
                </div>
                
                <div class="col-md-4">
                    <div class="card mb-4 box-shadow">
                        <a href="/product/66VCHSJNUP">
                            <img class="card-img-top" alt =""
                                style="width: 100%; height: auto;"
                                src="/static/img/products/camera-lens.jpg">
                        </a>
                        <div class="card-body">
                            <h5 class="card-title">
                                Vintage Camera Lens
                            </h5>
                            <div class="d-flex justify-content-between align-items-center">
                                <div class="btn-group">
                                    <a href="/product/66VCHSJNUP">
                                        <button type="button" class="btn btn-sm btn-outline-secondary">Buy</button>
                                    </a>
                                </div>
                                <small class="text-muted">
                                    USD 12.48 
                                </strong>
                                </small>
                            </div>
                        </div>
                    </div>
                </div>
                
                <div class="col-md-4">
                    <div class="card mb-4 box-shadow">
                        <a href="/product/1YMWWN1N4O">
                            <img class="card-img-top" alt =""
                                style="width: 100%; height: auto;"
                                src="/static/img/products/barista-kit.jpg">
                        </a>
                        <div class="card-body">
                            <h5 class="card-title">
                                Home Barista Kit
                            </h5>
                            <div class="d-flex justify-content-between align-items-center">
                                <div class="btn-group">
                                    <a href="/product/1YMWWN1N4O">
                                        <button type="button" class="btn btn-sm btn-outline-secondary">Buy</button>
                                    </a>
                                </div>
                                <small class="text-muted">
                                    USD 123.99 
                                </strong>
                                </small>
                            </div>
                        </div>
                    </div>
                </div>
                
                <div class="col-md-4">
                    <div class="card mb-4 box-shadow">
                        <a href="/product/L9ECAV7KIM">
                            <img class="card-img-top" alt =""
                                style="width: 100%; height: auto;"
                                src="/static/img/products/terrarium.jpg">
                        </a>
                        <div class="card-body">
                            <h5 class="card-title">
                                Terrarium
                            </h5>
                            <div class="d-flex justify-content-between align-items-center">
                                <div class="btn-group">
                                    <a href="/product/L9ECAV7KIM">
                                        <button type="button" class="btn btn-sm btn-outline-secondary">Buy</button>
                                    </a>
                                </div>
                                <small class="text-muted">
                                    USD 36.45 
                                </strong>
                                </small>
                            </div>
                        </div>
                    </div>
                </div>
                
                <div class="col-md-4">
                    <div class="card mb-4 box-shadow">
                        <a href="/product/2ZYFJ3GM2N">
                            <img class="card-img-top" alt =""
                                style="width: 100%; height: auto;"
                                src="/static/img/products/film-camera.jpg">
                        </a>
                        <div class="card-body">
                            <h5 class="card-title">
                                Film Camera
                            </h5>
                            <div class="d-flex justify-content-between align-items-center">
                                <div class="btn-group">
                                    <a href="/product/2ZYFJ3GM2N">
                                        <button type="button" class="btn btn-sm btn-outline-secondary">Buy</button>
                                    </a>
                                </div>
                                <small class="text-muted">
                                    USD 2244.99 
                                </strong>
                                </small>
                            </div>
                        </div>
                    </div>
                </div>
                
                <div class="col-md-4">
                    <div class="card mb-4 box-shadow">
                        <a href="/product/0PUK6V6EV0">
                            <img class="card-img-top" alt =""
                                style="width: 100%; height: auto;"
                                src="/static/img/products/record-player.jpg">
                        </a>
                        <div class="card-body">
                            <h5 class="card-title">
                                Vintage Record Player
                            </h5>
                            <div class="d-flex justify-content-between align-items-center">
                                <div class="btn-group">
                                    <a href="/product/0PUK6V6EV0">
                                        <button type="button" class="btn btn-sm btn-outline-secondary">Buy</button>
                                    </a>
                                </div>
                                <small class="text-muted">
                                    USD 65.49 
                                </strong>
                                </small>
                            </div>
                        </div>
                    </div>
                </div>
                
                <div class="col-md-4">
                    <div class="card mb-4 box-shadow">
                        <a href="/product/LS4PSXUNUM">
                            <img class="card-img-top" alt =""
                                style="width: 100%; height: auto;"
                                src="/static/img/products/camp-mug.jpg">
                        </a>
                        <div class="card-body">
                            <h5 class="card-title">
                                Metal Camping Mug
                            </h5>
                            <div class="d-flex justify-content-between align-items-center">
                                <div class="btn-group">
                                    <a href="/product/LS4PSXUNUM">
                                        <button type="button" class="btn btn-sm btn-outline-secondary">Buy</button>
                                    </a>
                                </div>
                                <small class="text-muted">
                                    USD 24.32 
                                </strong>
                                </small>
                            </div>
                        </div>
                    </div>
                </div>
                
                <div class="col-md-4">
                    <div class="card mb-4 box-shadow">
                        <a href="/product/9SIQT8TOJO">
                            <img class="card-img-top" alt =""
                                style="width: 100%; height: auto;"
                                src="/static/img/products/city-bike.jpg">
                        </a>
                        <div class="card-body">
                            <h5 class="card-title">
                                City Bike
                            </h5>
                            <div class="d-flex justify-content-between align-items-center">
                                <div class="btn-group">
                                    <a href="/product/9SIQT8TOJO">
                                        <button type="button" class="btn btn-sm btn-outline-secondary">Buy</button>
                                    </a>
                                </div>
                                <small class="text-muted">
                                    USD 789.49 
                                </strong>
                                </small>
                            </div>
                        </div>
                    </div>
                </div>
                
                <div class="col-md-4">
                    <div class="card mb-4 box-shadow">
                        <a href="/product/6E92ZMYYFZ">
                            <img class="card-img-top" alt =""
                                style="width: 100%; height: auto;"
                                src="/static/img/products/air-plant.jpg">
                        </a>
                        <div class="card-body">
                            <h5 class="card-title">
                                Air Plant
                            </h5>
                            <div class="d-flex justify-content-between align-items-center">
                                <div class="btn-group">
                                    <a href="/product/6E92ZMYYFZ">
                                        <button type="button" class="btn btn-sm btn-outline-secondary">Buy</button>
                                    </a>
                                </div>
                                <small class="text-muted">
                                    USD 12.29 
                                </strong>
                                </small>
                            </div>
                        </div>
                    </div>
                </div>
                
            </div>
            <div class="row">
                
<div class="container">
    <div class="alert alert-dark" role="alert">
        <strong>Advertisement:</strong>
        <a href="https://www.google.com" rel="nofollow" target="_blank" class="alert-link">
            default product, price: 0 USD
        </a>
    </div>
</div>

            </div>

            
                <div class="trace">
{
	&#34;id&#34;: 1,
	&#34;records&#34;: [
		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589923945754595459, &#34;uuid&#34;: &#34;generated by client&#34;, messageName&#34;: &#34;/&#34;},
		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589923945754963287, &#34;uuid&#34;: &#34;6bc744a9-8c97-8266-89bd-9f0083d4d0b5&#34;, &#34;messageName&#34;: &#34;ListProductsRequest&#34;},
		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589923945757280363, &#34;uuid&#34;: &#34;6bc744a9-8c97-8266-89bd-9f0083d4d0b5&#34;, &#34;messageName&#34;: &#34;ListProductsRequest&#34;},
		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589923945759155010, &#34;uuid&#34;: &#34;6bc744a9-8c97-8266-89bd-9f0083d4d0b5&#34;, &#34;messageName&#34;: &#34;ListProductsResponse&#34;},
		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589923945825398593, &#34;uuid&#34;: &#34;6bc744a9-8c97-8266-89bd-9f0083d4d0b5&#34;, &#34;messageName&#34;: &#34;ListProductsResponse&#34;},
		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589923951345510487, &#34;uuid&#34;: &#34;a50b3122-f3f2-f1c0-ef25-6b07734133ba&#34;, &#34;messageName&#34;: &#34;GetSupportedCurrenciesRequest&#34;},
		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589923951345510487, &#34;uuid&#34;: &#34;a50b3122-f3f2-f1c0-ef25-6b07734133ba&#34;, &#34;messageName&#34;: &#34;GetSupportedCurrenciesRequest&#34;},
		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589923952480000000, &#34;uuid&#34;: &#34;a50b3122-f3f2-f1c0-ef25-6b07734133ba&#34;, &#34;messageName&#34;: &#34;&#34;},
		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589923952484171159, &#34;uuid&#34;: &#34;a50b3122-f3f2-f1c0-ef25-6b07734133ba&#34;, &#34;messageName&#34;: &#34;GetSupportedCurrenciesResponse&#34;},
		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589923952484447835, &#34;uuid&#34;: &#34;02b68850-4b0d-5021-3beb-518520b4570e&#34;, &#34;messageName&#34;: &#34;ListProductsRequest&#34;},
		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589923952484791454, &#34;uuid&#34;: &#34;02b68850-4b0d-5021-3beb-518520b4570e&#34;, &#34;messageName&#34;: &#34;ListProductsRequest&#34;},
		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589923952485389250, &#34;uuid&#34;: &#34;02b68850-4b0d-5021-3beb-518520b4570e&#34;, &#34;messageName&#34;: &#34;ListProductsResponse&#34;},
		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589923952485861685, &#34;uuid&#34;: &#34;02b68850-4b0d-5021-3beb-518520b4570e&#34;, &#34;messageName&#34;: &#34;ListProductsResponse&#34;},
		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589923952496106038, &#34;uuid&#34;: &#34;23976a46-03a4-3334-c714-ec8b7f166d17&#34;, &#34;messageName&#34;: &#34;GetCartRequest&#34;},
		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589923952499693853, &#34;uuid&#34;: &#34;d403df70-f49f-077e-953d-1aa268455230&#34;, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589923952499693853, &#34;uuid&#34;: &#34;d403df70-f49f-077e-953d-1aa268455230&#34;, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589923952503000064, &#34;uuid&#34;: &#34;d403df70-f49f-077e-953d-1aa268455230&#34;, &#34;messageName&#34;: &#34;&#34;},
		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589923952504703538, &#34;uuid&#34;: &#34;d403df70-f49f-077e-953d-1aa268455230&#34;, &#34;messageName&#34;: &#34;Money&#34;},
		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589923952504771255, &#34;uuid&#34;: &#34;ab691fd2-e5ec-7880-d4fc-dee0f54bd4d3&#34;, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589923952504771255, &#34;uuid&#34;: &#34;ab691fd2-e5ec-7880-d4fc-dee0f54bd4d3&#34;, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589923952504999936, &#34;uuid&#34;: &#34;ab691fd2-e5ec-7880-d4fc-dee0f54bd4d3&#34;, &#34;messageName&#34;: &#34;&#34;},
		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589923952505990819, &#34;uuid&#34;: &#34;ab691fd2-e5ec-7880-d4fc-dee0f54bd4d3&#34;, &#34;messageName&#34;: &#34;Money&#34;},
		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589923952506077718, &#34;uuid&#34;: &#34;4d923ac7-1cad-10a2-4a9c-bcc1480da0d6&#34;, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589923952506077718, &#34;uuid&#34;: &#34;4d923ac7-1cad-10a2-4a9c-bcc1480da0d6&#34;, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589923952505999872, &#34;uuid&#34;: &#34;4d923ac7-1cad-10a2-4a9c-bcc1480da0d6&#34;, &#34;messageName&#34;: &#34;&#34;},
		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589923952507115744, &#34;uuid&#34;: &#34;4d923ac7-1cad-10a2-4a9c-bcc1480da0d6&#34;, &#34;messageName&#34;: &#34;Money&#34;},
		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589923952507321167, &#34;uuid&#34;: &#34;d20b553f-685f-46df-b551-9dc7ac528b37&#34;, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589923952507321167, &#34;uuid&#34;: &#34;d20b553f-685f-46df-b551-9dc7ac528b37&#34;, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589923952508000000, &#34;uuid&#34;: &#34;d20b553f-685f-46df-b551-9dc7ac528b37&#34;, &#34;messageName&#34;: &#34;&#34;},
		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589923952508470515, &#34;uuid&#34;: &#34;d20b553f-685f-46df-b551-9dc7ac528b37&#34;, &#34;messageName&#34;: &#34;Money&#34;},
		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589923952508628392, &#34;uuid&#34;: &#34;372f4217-be7d-efb2-b786-2a08c61edcc9&#34;, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589923952508628392, &#34;uuid&#34;: &#34;372f4217-be7d-efb2-b786-2a08c61edcc9&#34;, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589923952508999936, &#34;uuid&#34;: &#34;372f4217-be7d-efb2-b786-2a08c61edcc9&#34;, &#34;messageName&#34;: &#34;&#34;},
		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589923952509588523, &#34;uuid&#34;: &#34;372f4217-be7d-efb2-b786-2a08c61edcc9&#34;, &#34;messageName&#34;: &#34;Money&#34;},
		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589923952509671816, &#34;uuid&#34;: &#34;1a60e3c0-eda0-7db2-4eee-ef19c63aaf86&#34;, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589923952509671816, &#34;uuid&#34;: &#34;1a60e3c0-eda0-7db2-4eee-ef19c63aaf86&#34;, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589923952510000128, &#34;uuid&#34;: &#34;1a60e3c0-eda0-7db2-4eee-ef19c63aaf86&#34;, &#34;messageName&#34;: &#34;&#34;},
		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589923952510563506, &#34;uuid&#34;: &#34;1a60e3c0-eda0-7db2-4eee-ef19c63aaf86&#34;, &#34;messageName&#34;: &#34;Money&#34;},
		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589923952510596336, &#34;uuid&#34;: &#34;44666c09-e052-a1c4-880c-5e6164c45195&#34;, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589923952510596336, &#34;uuid&#34;: &#34;44666c09-e052-a1c4-880c-5e6164c45195&#34;, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589923952511000064, &#34;uuid&#34;: &#34;44666c09-e052-a1c4-880c-5e6164c45195&#34;, &#34;messageName&#34;: &#34;&#34;},
		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589923952512499996, &#34;uuid&#34;: &#34;44666c09-e052-a1c4-880c-5e6164c45195&#34;, &#34;messageName&#34;: &#34;Money&#34;},
		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589923952512559381, &#34;uuid&#34;: &#34;793a10e6-7185-2604-fff0-aa1bef6b7ce4&#34;, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589923952512559381, &#34;uuid&#34;: &#34;793a10e6-7185-2604-fff0-aa1bef6b7ce4&#34;, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589923952512999936, &#34;uuid&#34;: &#34;793a10e6-7185-2604-fff0-aa1bef6b7ce4&#34;, &#34;messageName&#34;: &#34;&#34;},
		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589923952513633051, &#34;uuid&#34;: &#34;793a10e6-7185-2604-fff0-aa1bef6b7ce4&#34;, &#34;messageName&#34;: &#34;Money&#34;},
		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589923952513729013, &#34;uuid&#34;: &#34;65726587-24a1-c99a-39c3-011f70a03823&#34;, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589923952513729013, &#34;uuid&#34;: &#34;65726587-24a1-c99a-39c3-011f70a03823&#34;, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589923952513999872, &#34;uuid&#34;: &#34;65726587-24a1-c99a-39c3-011f70a03823&#34;, &#34;messageName&#34;: &#34;&#34;},
		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589923952514810013, &#34;uuid&#34;: &#34;65726587-24a1-c99a-39c3-011f70a03823&#34;, &#34;messageName&#34;: &#34;Money&#34;},
		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589923952514974360, &#34;uuid&#34;: &#34;15d68f2a-4101-6ec1-1c06-7cfc58b7ef47&#34;, &#34;messageName&#34;: &#34;AdRequest&#34;},
		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589923952615372820, &#34;uuid&#34;: &#34;f8ed2a96-e340-da62-de21-74f8aa10d5c5&#34;, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589923952615372820, &#34;uuid&#34;: &#34;f8ed2a96-e340-da62-de21-74f8aa10d5c5&#34;, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589923952616000000, &#34;uuid&#34;: &#34;f8ed2a96-e340-da62-de21-74f8aa10d5c5&#34;, &#34;messageName&#34;: &#34;&#34;},
		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589923952617533676, &#34;uuid&#34;: &#34;f8ed2a96-e340-da62-de21-74f8aa10d5c5&#34;, &#34;messageName&#34;: &#34;Money&#34;},
		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589923952617566808, &#34;uuid&#34;: &#34;generated by client&#34;, &#34;messageName&#34;: &#34;/&#34;}
	],
	&#34;rlfi&#34;: [
		null
	],
	&#34;tfi&#34;: [
		null
	]
}
                </div>
            
            </div>
        </div>
    </main>

    
    <footer class="py-5 px-5">
        <div class="container">
            <p>
                &copy; 2018 Google Inc
                <span class="text-muted">
                    <a href="https://github.com/AleckDarcy/opencensus-microservices-demo/">(Source Code)</a>
                </span>
            </p>
            <p>
                <small class="text-muted">
                    This website is hosted for demo purposes only. It is not an
                    actual shop. This is not an official Google project.
                </small>
            </p>
            <small class="text-muted">
                session-id: b71503e8-b487-4f12-b664-4174accd5c61</br>
                request-id: bc1774f8-0d4f-44db-95e2-5a2c2cc08288</br>
            </small>
        </div>
    </footer>
    <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.1.1/js/bootstrap.min.js" integrity="sha384-smHYKdLADwkXOn1EmN1qk/HfnUcbVRZyYmZ4qpPea6sjB/pTJ0euyQp0Mk8ck+5T" crossorigin="anonymous"></script>
</body>
</html>
`

func TestHTML(t *testing.T) {
	node, _ := html.Parse(strings.NewReader(str))

	trace := GetElementByClass(node, "trace")

	t.Log(GetJSON(trace))
}

func BenchmarkName(b *testing.B) {
	trace := &tracer.Trace{}
	traceStr := `{
        	"id": 1,
        	"records": [
        		{"type": 2, "timestamp": 1589923945754595459, "uuid": "generated by client", "messageName": "/"},
        		{"type": 1, "timestamp": 1589923945754963287, "uuid": "6bc744a9-8c97-8266-89bd-9f0083d4d0b5", "messageName": "ListProductsRequest"},
        		{"type": 2, "timestamp": 1589923945757280363, "uuid": "6bc744a9-8c97-8266-89bd-9f0083d4d0b5", "messageName": "ListProductsRequest"},
        		{"type": 1, "timestamp": 1589923945759155010, "uuid": "6bc744a9-8c97-8266-89bd-9f0083d4d0b5", "messageName": "ListProductsResponse"},
        		{"type": 2, "timestamp": 1589923945825398593, "uuid": "6bc744a9-8c97-8266-89bd-9f0083d4d0b5", "messageName": "ListProductsResponse"},
        		{"type": 1, "timestamp": 1589923951345510487, "uuid": "a50b3122-f3f2-f1c0-ef25-6b07734133ba", "messageName": "GetSupportedCurrenciesRequest"},
        		{"type": 2, "timestamp": 1589923951345510487, "uuid": "a50b3122-f3f2-f1c0-ef25-6b07734133ba", "messageName": "GetSupportedCurrenciesRequest"},
        		{"type": 1, "timestamp": 1589923952480000000, "uuid": "a50b3122-f3f2-f1c0-ef25-6b07734133ba", "messageName": ""},
        		{"type": 2, "timestamp": 1589923952484171159, "uuid": "a50b3122-f3f2-f1c0-ef25-6b07734133ba", "messageName": "GetSupportedCurrenciesResponse"},
        		{"type": 1, "timestamp": 1589923952484447835, "uuid": "02b68850-4b0d-5021-3beb-518520b4570e", "messageName": "ListProductsRequest"},
        		{"type": 2, "timestamp": 1589923952484791454, "uuid": "02b68850-4b0d-5021-3beb-518520b4570e", "messageName": "ListProductsRequest"},
        		{"type": 1, "timestamp": 1589923952485389250, "uuid": "02b68850-4b0d-5021-3beb-518520b4570e", "messageName": "ListProductsResponse"},
        		{"type": 2, "timestamp": 1589923952485861685, "uuid": "02b68850-4b0d-5021-3beb-518520b4570e", "messageName": "ListProductsResponse"},
        		{"type": 1, "timestamp": 1589923952496106038, "uuid": "23976a46-03a4-3334-c714-ec8b7f166d17", "messageName": "GetCartRequest"},
        		{"type": 1, "timestamp": 1589923952499693853, "uuid": "d403df70-f49f-077e-953d-1aa268455230", "messageName": "CurrencyConversionRequest"},
        		{"type": 2, "timestamp": 1589923952499693853, "uuid": "d403df70-f49f-077e-953d-1aa268455230", "messageName": "CurrencyConversionRequest"},
        		{"type": 1, "timestamp": 1589923952503000064, "uuid": "d403df70-f49f-077e-953d-1aa268455230", "messageName": ""},
        		{"type": 2, "timestamp": 1589923952504703538, "uuid": "d403df70-f49f-077e-953d-1aa268455230", "messageName": "Money"},
        		{"type": 1, "timestamp": 1589923952504771255, "uuid": "ab691fd2-e5ec-7880-d4fc-dee0f54bd4d3", "messageName": "CurrencyConversionRequest"},
        		{"type": 2, "timestamp": 1589923952504771255, "uuid": "ab691fd2-e5ec-7880-d4fc-dee0f54bd4d3", "messageName": "CurrencyConversionRequest"},
        		{"type": 1, "timestamp": 1589923952504999936, "uuid": "ab691fd2-e5ec-7880-d4fc-dee0f54bd4d3", "messageName": ""},
        		{"type": 2, "timestamp": 1589923952505990819, "uuid": "ab691fd2-e5ec-7880-d4fc-dee0f54bd4d3", "messageName": "Money"},
        		{"type": 1, "timestamp": 1589923952506077718, "uuid": "4d923ac7-1cad-10a2-4a9c-bcc1480da0d6", "messageName": "CurrencyConversionRequest"},
        		{"type": 2, "timestamp": 1589923952506077718, "uuid": "4d923ac7-1cad-10a2-4a9c-bcc1480da0d6", "messageName": "CurrencyConversionRequest"},
        		{"type": 1, "timestamp": 1589923952505999872, "uuid": "4d923ac7-1cad-10a2-4a9c-bcc1480da0d6", "messageName": ""},
        		{"type": 2, "timestamp": 1589923952507115744, "uuid": "4d923ac7-1cad-10a2-4a9c-bcc1480da0d6", "messageName": "Money"},
        		{"type": 1, "timestamp": 1589923952507321167, "uuid": "d20b553f-685f-46df-b551-9dc7ac528b37", "messageName": "CurrencyConversionRequest"},
        		{"type": 2, "timestamp": 1589923952507321167, "uuid": "d20b553f-685f-46df-b551-9dc7ac528b37", "messageName": "CurrencyConversionRequest"},
        		{"type": 1, "timestamp": 1589923952508000000, "uuid": "d20b553f-685f-46df-b551-9dc7ac528b37", "messageName": ""},
        		{"type": 2, "timestamp": 1589923952508470515, "uuid": "d20b553f-685f-46df-b551-9dc7ac528b37", "messageName": "Money"},
        		{"type": 1, "timestamp": 1589923952508628392, "uuid": "372f4217-be7d-efb2-b786-2a08c61edcc9", "messageName": "CurrencyConversionRequest"},
        		{"type": 2, "timestamp": 1589923952508628392, "uuid": "372f4217-be7d-efb2-b786-2a08c61edcc9", "messageName": "CurrencyConversionRequest"},
        		{"type": 1, "timestamp": 1589923952508999936, "uuid": "372f4217-be7d-efb2-b786-2a08c61edcc9", "messageName": ""},
        		{"type": 2, "timestamp": 1589923952509588523, "uuid": "372f4217-be7d-efb2-b786-2a08c61edcc9", "messageName": "Money"},
        		{"type": 1, "timestamp": 1589923952509671816, "uuid": "1a60e3c0-eda0-7db2-4eee-ef19c63aaf86", "messageName": "CurrencyConversionRequest"},
        		{"type": 2, "timestamp": 1589923952509671816, "uuid": "1a60e3c0-eda0-7db2-4eee-ef19c63aaf86", "messageName": "CurrencyConversionRequest"},
        		{"type": 1, "timestamp": 1589923952510000128, "uuid": "1a60e3c0-eda0-7db2-4eee-ef19c63aaf86", "messageName": ""},
        		{"type": 2, "timestamp": 1589923952510563506, "uuid": "1a60e3c0-eda0-7db2-4eee-ef19c63aaf86", "messageName": "Money"},
        		{"type": 1, "timestamp": 1589923952510596336, "uuid": "44666c09-e052-a1c4-880c-5e6164c45195", "messageName": "CurrencyConversionRequest"},
        		{"type": 2, "timestamp": 1589923952510596336, "uuid": "44666c09-e052-a1c4-880c-5e6164c45195", "messageName": "CurrencyConversionRequest"},
        		{"type": 1, "timestamp": 1589923952511000064, "uuid": "44666c09-e052-a1c4-880c-5e6164c45195", "messageName": ""},
        		{"type": 2, "timestamp": 1589923952512499996, "uuid": "44666c09-e052-a1c4-880c-5e6164c45195", "messageName": "Money"},
        		{"type": 1, "timestamp": 1589923952512559381, "uuid": "793a10e6-7185-2604-fff0-aa1bef6b7ce4", "messageName": "CurrencyConversionRequest"},
        		{"type": 2, "timestamp": 1589923952512559381, "uuid": "793a10e6-7185-2604-fff0-aa1bef6b7ce4", "messageName": "CurrencyConversionRequest"},
        		{"type": 1, "timestamp": 1589923952512999936, "uuid": "793a10e6-7185-2604-fff0-aa1bef6b7ce4", "messageName": ""},
        		{"type": 2, "timestamp": 1589923952513633051, "uuid": "793a10e6-7185-2604-fff0-aa1bef6b7ce4", "messageName": "Money"},
        		{"type": 1, "timestamp": 1589923952513729013, "uuid": "65726587-24a1-c99a-39c3-011f70a03823", "messageName": "CurrencyConversionRequest"},
        		{"type": 2, "timestamp": 1589923952513729013, "uuid": "65726587-24a1-c99a-39c3-011f70a03823", "messageName": "CurrencyConversionRequest"},
        		{"type": 1, "timestamp": 1589923952513999872, "uuid": "65726587-24a1-c99a-39c3-011f70a03823", "messageName": ""},
        		{"type": 2, "timestamp": 1589923952514810013, "uuid": "65726587-24a1-c99a-39c3-011f70a03823", "messageName": "Money"},
        		{"type": 1, "timestamp": 1589923952514974360, "uuid": "15d68f2a-4101-6ec1-1c06-7cfc58b7ef47", "messageName": "AdRequest"},
        		{"type": 1, "timestamp": 1589923952615372820, "uuid": "f8ed2a96-e340-da62-de21-74f8aa10d5c5", "messageName": "CurrencyConversionRequest"},
        		{"type": 2, "timestamp": 1589923952615372820, "uuid": "f8ed2a96-e340-da62-de21-74f8aa10d5c5", "messageName": "CurrencyConversionRequest"},
        		{"type": 1, "timestamp": 1589923952616000000, "uuid": "f8ed2a96-e340-da62-de21-74f8aa10d5c5", "messageName": ""},
        		{"type": 2, "timestamp": 1589923952617533676, "uuid": "f8ed2a96-e340-da62-de21-74f8aa10d5c5", "messageName": "Money"},
        		{"type": 1, "timestamp": 1589923952617566808, "uuid": "generated by client", "messageName": "/"}
        	],
        	"rlfi": [
        		null
        	],
        	"tfi": [
        		null
        	]
        }`
	json.Unmarshal([]byte(traceStr), trace)

	traceBytes := []byte{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		traceBytes, _ = json.Marshal(trace)
	}

	b.Log(traceStr)
	b.Log(traceBytes)
}
