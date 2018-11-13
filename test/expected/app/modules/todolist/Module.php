<?php

namespace app\modules\todolist;

use Yii;
use yii\web\Response;
use yii\base\BootstrapInterface;

/**
 * api module definition class
 */
class Module extends \yii\base\Module implements BootstrapInterface
{
    /**
     * {@inheritdoc}
     */
    public $controllerNamespace = 'app\modules\todolist\controllers';

    /**
     * {@inheritdoc}
     */
    public function init()
    {
        parent::init();
        Yii::$app->response->format = Response::FORMAT_JSON;

        Yii::$app->setComponents([
            'request' => [
                'class' => \yii\web\Request::class,
                'parsers' => [
                    'application/json' => 'yii\web\JsonParser',
                ],
                'enableCookieValidation' => false,
                'enableCsrfValidation' => false,
            ],
            'errorHandler' => [
                'class' => 'app\modules\todolist\handlers\ErrorHandler',
            ],
        ]);

        $handler = $this->get('errorHandler');
        \Yii::$app->set('errorHandler', $handler);
        $handler->register();
    }

    public function bootstrap($app)
    {
        $app->getUrlManager()->addRules([
            "POST TodolistService.add" => "TodolistService/api/add",
            "POST TodolistService.list" => "TodolistService/api/list",
        ], false);
    }
}
