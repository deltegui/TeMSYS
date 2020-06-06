<?php
require '../src/user/user.controller.php';

use Temsys\Controllers\UserController;

$userController = new UserController();

$app->get("/", [$userController, '_helloIndex']);
