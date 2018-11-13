<?php
namespace app\modules\todolist\models;

use Yoozoo\ProtoApi;

class Todo implements ProtoApi\Message
{
    protected $title;

    public function init(array $response)
    {
        if (isset($response["title"])) {
            $this->title = $response["title"];
        }
    }

    public function validate()
    {
        if (!isset($this->title)) {
            throw new ProtoApi\GeneralException("'title' is not exist");
        }
    }
    
    public function set_title($title)
    {
        $this->title = $title;
    }

    public function get_title()
    {
        return $this->title;
    }
    
    public function to_array()
    {
        return array(
            "title" => $this->title,
        );
    }
}
