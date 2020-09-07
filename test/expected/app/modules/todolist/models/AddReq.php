<?php
namespace app\modules\todolist\models;

use Yoozoo\ProtoApi;

class AddReq implements ProtoApi\Message
{
    protected $item;

    public function init(array $response)
    {
        if (isset($response["item"])) {
            $this->item = new Todo();
            $this->item->init($response["item"]);
        }
    }

    public function validate()
    {
        if (!isset($this->item)) {
            throw new ProtoApi\GeneralException("'item' is not exist");
        }
        $this->item->validate();
    }
    
    public function set_item(Todo $item)
    {
        $this->item = $item;
    }

    public function get_item()
    {
        return $this->item;
    }
    
    public function to_array()
    {
        return array(
            "item" =>  $this->item->to_array(),
        );
    }
}
