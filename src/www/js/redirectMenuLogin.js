function redirectMenuLogin(id ){
    let httpRedirect = "";
    switch (id){
        case 1:
            httpRedirect = "/auth/my_board";
            break;
        case 3:
            httpRedirect = "/auth/create_guild";
            break;
        case 0:
            httpRedirect = "/logout"
    }
    window.location.href = httpRedirect;
}