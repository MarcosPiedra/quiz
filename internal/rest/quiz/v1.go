package quiz

import (
	"encoding/json"
	"net/http"
	"quiz-system/internal/application"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator"
	"github.com/stackus/errors"
)

type QuizV1 struct {
	app application.App
}

func NewQuizV1(app application.App) *QuizV1 {
	return &QuizV1{
		app: app,
	}
}

func (c *QuizV1) Register(r chi.Router) {
	r.Get("/questions", c.getQuestions)
	r.Get("/questions/{id}", c.getQuestion)
	r.Post("/evaluation", c.addEvaluation)
}

// getQuestions godoc
// @Summary      Get questions
// @Description  Get questions
// @Tags         Questions
// @Produce      json
// @Success      200  {object}  quiz.QuestionsResponse
// @Failure      500  {object}  quiz.ErrorResponse
// @Router       /quiz/v1/questions [get]
func (h *QuizV1) getQuestions(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	questions, err := h.app.GetQuestions(ctx)

	if questions == nil {
		renderNotFound(w, r)

		return
	}

	if err == nil {
		err = render.Render(w, r, NewQuestionsResponse(questions))
	}

	if err != nil {
		renderError(err, w, r)
	}
}

// getQuestion godoc
// @Summary      Get question
// @Description  Get question
// @Tags         Questions
// @Produce      json
// @Param id path int true "Question Id"
// @Success      200  {object}  quiz.QuestionResponse
// @Failure      404  {object}  quiz.ErrorResponse
// @Failure      500  {object}  quiz.ErrorResponse
// @Router       /quiz/v1/questions/{id} [get]
func (h *QuizV1) getQuestion(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		renderError(err, w, r)

		return
	}

	ctx := r.Context()

	question, err := h.app.GetQuestion(id, ctx)

	if question == nil {
		renderNotFound(w, r)

		return
	}

	if err == nil {
		render.Render(w, r, NewQuestionResponse(question))
	} else {
		renderError(err, w, r)
	}
}

// addEvaluation godoc
// @Summary      New evaluation
// @Description  New evaluation
// @Tags         Evaluation
// @Param        request    body    quiz.EvaluationsRequest    true    "Evaluation request"
// @Produce      json
// @Success      200  {object}  quiz.EvaluationResponse
// @Failure      500  {object}  quiz.ErrorResponse
// @Router       /quiz/v1/evaluation [post]
func (h *QuizV1) addEvaluation(w http.ResponseWriter, r *http.Request) {
	var evaluationRequest EvaluationsRequest

	err := json.NewDecoder(r.Body).Decode(&evaluationRequest)
	if err != nil {
		renderError(err, w, r)

		return
	}

	v := validator.New()
	err = v.Struct(evaluationRequest)
	if err != nil {
		renderBadRequest(err, w, r)

		return
	}

	addEvaluationCommand := evaluationRequest.ToCommand()
	evaluation, err := h.app.AddEvaluation(addEvaluationCommand, r.Context())

	if err == nil {
		render.Render(w, r, NewEvaluationResponse(evaluation))
	} else {
		renderError(err, w, r)
	}
}

func renderBadRequest(err error, w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, Err(
		errors.ErrBadRequest.Err(err),
		errors.ErrBadRequest.HTTPCode()))
}

func renderNotFound(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, Err(
		errors.ErrNotFound.Msg("Not found"),
		errors.ErrNotFound.HTTPCode()))
}

func renderError(err error, w http.ResponseWriter, r *http.Request) {
	var coder errors.HTTPCoder = nil
	if errors.As(err, &coder) {
		render.Render(w, r, Err(err, coder.HTTPCode()))
	} else {
		render.Render(w, r, Err(
			errors.ErrInternalServerError.Msg("Internal server error"),
			errors.ErrInternalServerError.HTTPCode()))
	}
}
