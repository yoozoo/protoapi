<?php

namespace app\modules\todolist\controllers;

use app\modules\todolist\models;
use Yii;
use yii\web\Controller;
use Yoozoo\ProtoApi;

class ApiController extends Controller
{
    private $_handler;

    public function init()
    {
        $this->_handler = new \app\modules\todolist\RequestHandler();
    }

    /**
     * {@inheritdoc}
     */
    public function behaviors()
    {
        $behaviors = parent::behaviors();
        if (class_exists("\\app\\modules\\todolist\\AuthHandler")){
            $behaviors['authenticator'] = [
                'class' => \app\modules\todolist\AuthHandler::className(),
            ];
        }
        if (class_exists("\\app\\modules\\todolist\\CorsHandler")){
            $behaviors['authenticator'] = [
                'class' => \app\modules\todolist\CorsHandler::className(),
            ];
        }

        return $behaviors;
    }
    
    public function actionAdd()
    {
        $req = Yii::$app->request;
        $request = new models\AddReq();
        $request->init($req->getBodyParams());
        $request->validate();
        $res = $this->_handler->add($request);
        if ($res instanceof models\AddResp) {
            $res->validate();
            return $res->to_array();
        }
        throw new ProtoApi\GeneralException("return type of 'add' incorrect.");
    }
    
    public function actionList()
    {
        $req = Yii::$app->request;
        $request = new models\Blank();
        $request->init($req->getBodyParams());
        $request->validate();
        $res = $this->_handler->list($request);
        if ($res instanceof models\ListResp) {
            $res->validate();
            return $res->to_array();
        }
        throw new ProtoApi\GeneralException("return type of 'list' incorrect.");
    }
    
}
