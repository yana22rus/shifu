function edit_value(clicked_id,value_edit){

    console.log(String(clicked_id));
    let btn_clear = document.getElementById("input_data").value = "";
    let btn_add = document.getElementById("input_data").value = value_edit;
    let xxx = document.getElementById(String("edit_data")).value = clicked_id;

}