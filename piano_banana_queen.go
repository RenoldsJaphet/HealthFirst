package healthcare_provider

import "fmt"

// HealthcareProvider struct to represent a healthcare provider
type HealthcareProvider struct {
	Name string
	Location string
	Services []string
}

// CreateNewHealthcareProvider creates a new healthcare provider with provided details
func CreateNewHealthcareProvider(name string, location string, services []string) *HealthcareProvider {
	return &HealthcareProvider{
		Name: name,
		Location: location,
		Services: services,
	}
}

// GetName returns the name of the healthcare provider
func (h *HealthcareProvider) GetName() string {
	return h.Name
}

// GetLocation returns the location of the healthcare provider
func (h *HealthcareProvider) GetLocation() string {
	return h.Location
}

// GetServices returns the services provided by the healthcare provider
func (h *HealthcareProvider) GetServices() []string {
	return h.Services
}

// SetName sets the name for the healthcare provider
func (h *HealthcareProvider) SetName(name string) {
	h.Name = name
}

// SetLocation sets the location of the healthcare provider
func (h *HealthcareProvider) SetLocation(location string) {
	h.Location = location
}

// SetServices sets the services for the healthcare provider
func (h *HealthcareProvider) SetServices(services []string) {
	h.Services = services
}

// PrintServices prints the services provided by the healthcare provider
func (h *HealthcareProvider) PrintServices() {
	fmt.Println("The services provided by this healthcare provider are:")
	for _, service := range h.Services {
		fmt.Println(service)
	}
}

// AddService adds a service to the healthcare provider
func (h *HealthcareProvider) AddService(service string) {
	h.Services = append(h.Services, service)
}

// RemoveService removes a service from the healthcare provider
func (h *HealthcareProvider) RemoveService(service string) {
	index := -1
	for i, s := range h.Services {
		if s == service {
			index = i
			break
		}
	}
	if index >= 0 {
		h.Services = append(h.Services[:index], h.Services[index+1:]...)
	}
}