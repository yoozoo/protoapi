<?php

namespace app\modules\todolist;

use app\modules\todolist\models;

class RequestHandler extends handlers\RequestHandler{
    /**
     * @param models\AddReq $req
     * @return models\AddResp
     */
    function add(models\AddReq $req) {
        // implement here
    }
    
    /**
     * @param models\Blank $req
     * @return models\ListResp
     */
    function list(models\Blank $req) {
        // implement here
    }
    
}
