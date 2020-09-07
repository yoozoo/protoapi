<?php
namespace app\modules\todolist\models;

use Yoozoo\ProtoApi;

class ListResp implements ProtoApi\Message
{
    protected $items;

    public function init(array $response)
    {
        if (isset($response["items"])) {
            $val = $response["items"];
            $this->set_items( array_map( function($v) { $tmp = new Todo(); $tmp->init($v); return $tmp; }, $val) );
        }
    }

    public function validate()
    {
        if (!isset($this->items)) {
            throw new ProtoApi\GeneralException("'items' is not exist");
        }
        array_filter($this->items, function($v) { $v->validate(); return false; });
    }
    
    public function set_items(array $items)
    {
        $this->items = $items;
    }

    public function get_items()
    {
        return $this->items;
    }
    
    public function to_array()
    {
        return array(
            "items" => array_map( function ($v) {  return $v->to_array(); }, $this->items),
        );
    }
}
