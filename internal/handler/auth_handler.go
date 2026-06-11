package handler

import (
	"log"
	"net/http"
	"strings"
	"sydesk/pkg/domain"

	"github.com/gin-gonic/gin"
)

// AuthHandler define la estructura para los endpoints de autenticación
type AuthHandler struct {
	Service domain.AuthServices
}

// Login es el handler para el endpoint POST /login
func (h *AuthHandler) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Validar formato del JSON de entrada
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("error: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "El mensaje no pudo ser deserializado"})
		return
	}

	// Llamar al servicio de autenticación
	token, err := h.Service.Login(input.Email, input.Password)
	if err != nil {
		// Error genérico para no dar pistas a atacantes
		log.Println("error: ", err.Error())

		if err.Error() == "Usuario no ha sido verificado" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Esta cuenta aún no ha sido verificada"})
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// Auth es el Middleware que bloquea el acceso a rutas protegidas
func Auth(authService domain.AuthServices) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No Authorization header"})
			c.Abort()
			return
		}

		// Limpiar el prefijo Bearer con espacio
		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))

		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato del token inválido (debe ser Bearer)"})
			c.Abort()
			return
		}

		// Delegar la validación al servicio implementado en la capa de dominio
		claims, err := authService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido o expirado"})
			c.Abort()
			return
		}

		// Inyectar datos del usuario en el contexto de la petición
		// Nota: Asegúrate de que tu struct CustomClaims use 'UserId' (con U mayúscula)
		c.Set("user_id", claims.UserId)
		c.Set("email", claims.Email)

		c.Next()
	}
}
