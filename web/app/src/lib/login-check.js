const isTokenExpired = (token) => {
    if (token === "") return true;
    const tokenData = (JSON.parse(atob(token.split('.')[1])));
    const expiry = tokenData.exp;
    return (Math.floor((new Date).getTime() / 1000)) >= expiry;
}
const getUsernameFromToken = (token) => {
    if (token === "") return "logged out";

    const isExpired = isTokenExpired(token);

    if (!isExpired) {
        const tokenData = (JSON.parse(atob(token.split('.')[1])));
        return tokenData.name;
    }
    return "Guest";

}

export { isTokenExpired, getUsernameFromToken }