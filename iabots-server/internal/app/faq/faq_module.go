// package faq

// v1 := r.Group("/api/v1")

// bots := v1.Group("/bots")
// {
//     bots.POST("/:botId/faqs", faqHandler.Create)
//     bots.GET("/:botId/faqs", faqHandler.List)
//     bots.GET("/:botId/faqs/search", faqHandler.Search)
// }

// faqs := v1.Group("/faqs")
// {
//     faqs.GET("/:faqId", faqHandler.GetByID)
//     faqs.PUT("/:faqId", faqHandler.Update)
//     faqs.DELETE("/:faqId", faqHandler.Delete)
// }

package faq

import (
	. "iabots-server/internal/delivery/http/handlers"
	. "iabots-server/internal/domain/usecases/faq"
	. "iabots-server/internal/infra/database"
	. "iabots-server/internal/infra/database/repositories"
)

type FaqModule struct {
	Handler *FaqHandler
}

func NewFaqModule(db *Database) *FaqModule {
	/// Injection of binds - Faq Module
	/// --------------------------------------------------
	///  Repositories
	/// --------------------------------------------------
	faqRepo := NewFaqGormRepository(db.DB)
	botRepo := NewAssistantBotGormRepository(db.DB)
	/// --------------------------------------------------
	///  Use Cases
	/// --------------------------------------------------
	createFaqUseCase := NewCreateFaqUseCase(faqRepo, botRepo)
	/// --------------------------------------------------
	///  Handlers
	/// --------------------------------------------------
	handler := NewFaqHandler(createFaqUseCase)
	return &FaqModule{
		Handler: handler,
	}
}
