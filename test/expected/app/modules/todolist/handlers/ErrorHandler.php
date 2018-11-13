<?php

namespace app\modules\todolist\handlers;

use Yii;
use Yoozoo\ProtoApi;

class ErrorHandler extends \yii\base\ErrorHandler
{
    public function renderException($exception)
    {
        if ($exception instanceof ProtoApi\BizErrorException) {
            Yii::$app->response->statusCode = 400;
            $resp = $exception->to_array();
        } else if ($exception instanceof ProtoApi\CommonErrorException) {
            Yii::$app->response->statusCode = 420;
            $resp = $exception->to_array();
        } else {
            Yii::$app->response->statusCode = 500;
            $resp = array(
                "message"=>$exception->getMessage(),
                "stack"=>$exception->getTraceAsString(),
            );
        }
        Yii::$app->response->data = $resp;
        Yii::$app->response->send();
    }
}
