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

    public $prefix;
    public $module_name;
    /**
     * {@inheritdoc}
     */
    public function beforeAction ( $action )
    {
        if (!parent::beforeAction($action) )
        {
            return false;
        }
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
        Yii::$app->set('errorHandler', $handler);
        $handler->register();
        return true;
    }

    public function bootstrap($app)
    {

        $app->getUrlManager()->addRules([
            new \yii\web\GroupUrlRule ([
                'prefix' => $this->prefix ,
                'routePrefix' => $this->module_name,
                'rules' => [
                    'TodolistService.add' => 'api/add',
                    'TodolistService.list' => 'api/list',
                ],
            ])
        ], false);
    }
}
