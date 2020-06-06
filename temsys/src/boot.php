<?php

use Slim\Csrf\Guard;
use Slim\Factory\AppFactory;
use Twig\Environment;

$app = AppFactory::create();

$app->addRoutingMiddleware();
$app->addErrorMiddleware(true, true, true);

session_start();

$responseFactory = $app->getResponseFactory();
$csrf = new Guard($responseFactory);
$app->add($csrf);
\Temsys\Controllers\Controller::setCsrf($csrf);
