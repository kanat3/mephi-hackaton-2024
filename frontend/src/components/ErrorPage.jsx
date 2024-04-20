import styled from "@emotion/styled";

const Container = styled.div`
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 10px;
`;

export const ErrorComponent = ({ errorObject }) => {
    return (
        <Container>
            <span style={{ color: "white", fontSize: "18px" }}>Произошла ошибка, попробуйте перезагрузить страницу</span>
            <span style={{ color: "white", fontSize: "18px" }}>{errorObject ? errorObject.message : "Что-то пошло не так, проверьте подключение к интернету"}</span>
        </Container>
    )
}