<?php
namespace app\modules\todolist\handlers;

use app\modules\todolist\models;
use Yoozoo\ProtoApi;

abstract class RequestHandler
{
    abstract public function add(models\AddReq $req);
    abstract public function list(models\Blank $req);
}
