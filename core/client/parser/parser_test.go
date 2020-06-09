package parser

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestHTML(t *testing.T) {
	str := `
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
                                            USD 67.98 
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
                                            USD 12.49 
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
                                            USD 124.00 
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
                                            USD 65.50 
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
                    default product, price: 1 USD
                </a>
            </div>
        </div>
        
                    </div>
        
                    
                        <hr/>
		<div class="trace">
        {
        	&#34;id&#34;: 1589671029855950000,
        	&#34;records&#34;: [
        		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589671029865508760, &#34;messageName&#34;: &#34;/&#34;},
        		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589671029865557359, &#34;messageName&#34;: &#34;GetSupportedCurrenciesRequest&#34;},
        		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589671029865557359, &#34;messageName&#34;: &#34;GetSupportedCurrenciesRequest&#34;},
        		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589671029867000064, &#34;messageName&#34;: &#34;&#34;},
        		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589671029868529608, &#34;messageName&#34;: &#34;GetSupportedCurrenciesResponse&#34;},
        		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589671029868671071, &#34;messageName&#34;: &#34;ListProductsRequest&#34;},
        		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589671029869029438, &#34;messageName&#34;: &#34;ListProductsRequest&#34;},
        		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589671029869862837, &#34;messageName&#34;: &#34;ListProductsResponse&#34;},
        		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589671029870338345, &#34;messageName&#34;: &#34;ListProductsResponse&#34;},
        		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589671029870476884, &#34;messageName&#34;: &#34;GetCartRequest&#34;},
        		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589671029872900906, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
        		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589671029872900906, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
        		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589671029873999872, &#34;messageName&#34;: &#34;&#34;},
        		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589671029875114140, &#34;messageName&#34;: &#34;Money&#34;},
        		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589671029875181761, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
        		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589671029875181761, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
        		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589671029912000000, &#34;messageName&#34;: &#34;&#34;},
        		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589671029913485302, &#34;messageName&#34;: &#34;Money&#34;},
        		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589671029913589116, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
        		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589671029913589116, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
        		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589671029915000064, &#34;messageName&#34;: &#34;&#34;},
        		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589671029916588105, &#34;messageName&#34;: &#34;Money&#34;},
        		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589671029916642293, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
        		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589671029916642293, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
        		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589671029918000128, &#34;messageName&#34;: &#34;&#34;},
        		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589671029919682200, &#34;messageName&#34;: &#34;Money&#34;},
        		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589671029919823191, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
        		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589671029919823191, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
        		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589671029920999936, &#34;messageName&#34;: &#34;&#34;},
        		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589671029922611712, &#34;messageName&#34;: &#34;Money&#34;},
        		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589671029922740207, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
        		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589671029922740207, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
        		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589671029924000000, &#34;messageName&#34;: &#34;&#34;},
        		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589671029924952388, &#34;messageName&#34;: &#34;Money&#34;},
        		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589671029925018405, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
        		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589671029925018405, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
        		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589671029924999936, &#34;messageName&#34;: &#34;&#34;},
        		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589671029926614129, &#34;messageName&#34;: &#34;Money&#34;},
        		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589671029926693819, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
        		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589671029926693819, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
        		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589671029929999872, &#34;messageName&#34;: &#34;&#34;},
        		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589671029931798521, &#34;messageName&#34;: &#34;Money&#34;},
        		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589671029931908346, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
        		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589671029931908346, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
        		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589671029932999936, &#34;messageName&#34;: &#34;&#34;},
        		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589671029933859137, &#34;messageName&#34;: &#34;Money&#34;},
        		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589671029934035689, &#34;messageName&#34;: &#34;AdRequest&#34;},
        		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589671029934176368, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
        		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589671029934176368, &#34;messageName&#34;: &#34;CurrencyConversionRequest&#34;},
        		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589671029936000000, &#34;messageName&#34;: &#34;&#34;},
        		{&#34;type&#34;: 2, &#34;timestamp&#34;: 1589671029939741261, &#34;messageName&#34;: &#34;Money&#34;},
        		{&#34;type&#34;: 1, &#34;timestamp&#34;: 1589671029939780134, &#34;messageName&#34;: &#34;/&#34;}
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
                        session-id: 36758c83-42e9-48a7-ac6a-be8de155653e</br>
                        request-id: e7f26b26-d408-4cb9-81eb-3709c0f145aa</br>
                    </small>
                </div>
            </footer>
            <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.1.1/js/bootstrap.min.js" integrity="sha384-smHYKdLADwkXOn1EmN1qk/HfnUcbVRZyYmZ4qpPea6sjB/pTJ0euyQp0Mk8ck+5T" crossorigin="anonymous"></script>
        </body>
        </html>
`

	node, _ := html.Parse(strings.NewReader(str))

	trace := GetElementByClass(node, "trace")

	t.Log(GetJSON(trace))
}
