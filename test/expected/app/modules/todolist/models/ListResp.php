<?php
namespace app\modules\todolist\models;

use Yoozoo\ProtoApi;

class ListResp implements ProtoApi\Message
{
    protected $items;

    public function init(array $response)
    {
        if (isset($response["items"])) {
            $this->items = array();
            foreach ($response["items"] as $items) {
                $tmp = new Todo();
                $tmp->init($items);
                $tmp->validate();
                $this->items[] = $tmp;
            }
        }
    }

    public function validate()
    {
        if (!isset($this->items)) {
            throw new ProtoApi\GeneralException("'items' is not exist");
        }
    }
    
    public function set_items(Items $items)
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
            "items" => $this->items->to_array(),
        );
    }
}
