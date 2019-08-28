<?php

namespace app\modules\todolist\controllers;

use app\modules\todolist\models;
use Yii;
use yii\web\Controller;
use yii\base\InlineAction;
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
        // cors need to be the first filter so it will be effective in all request
        if (class_exists("\\app\\modules\\todolist\\CorsHandler")){
            return ["cors" => [ 'class' => \app\modules\todolist\CorsHandler::className() ]] + $behaviors;
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
    
    private $methods = [ 'add' => 'actionAdd',  'list' => 'actionList', ];

    public function createAction($id)
    {
        if (isset($this->methods[$id])) {
            return new InlineAction($id, $this, $this->methods[$id]);
        }
        return null;
    }
}
