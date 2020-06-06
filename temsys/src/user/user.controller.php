<?php

namespace Temsys\Controllers;

use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;

class UserController extends Controller {
    public function helloIndex() {
        $this->render('login/index.html', [
            'csrf' => $this->csrfData(),
            'imgNum' => rand(0, 2),
        ]);
    }
}
