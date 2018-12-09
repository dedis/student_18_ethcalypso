const LTSReply = document.getElementById("LTSReply");
const WriteAddress = document.getElementById("WriteAddress");
const ReadADdress = document.getElementById("ReadAddress");
const secret = document.getElementById("Secret");

$('#LTSIDform').submit(function(e){
    e.preventDefault();
    $.ajax({
        url:"/GetLTSID",
        type:'post',
        data: $('#LTSIDform').serialize(),
        success:function(data){
            //whatever you wanna do after the form is successfully submitted
            const d = JSON.stringify(data);
            console.log(typeof(d));
            LTSReply.innerHTML = d;
        }
    });
});

$('#WriteForm').submit(function(e){
    e.preventDefault();
    $.ajax({
        url:"/AddWrite",
        type:'post',
        data:$('#WriteForm').serialize(),
        success:function(data){
            //whatever you wanna do after the form is successfully submitted
            WriteAddress.innerHTML = data;
        }
    });
});

$('#ReadForm').submit(function(e){
    e.preventDefault();
    $.ajax({
        url:"/AddRead",
        type:'post',
        data:$('#ReadForm').serialize(),
        success:function(d){
            //whatever you wanna do after the form is successfully submitted
            ReadAddress.innerHTML = "Your read address and secret are " + d;
        }
    });
});