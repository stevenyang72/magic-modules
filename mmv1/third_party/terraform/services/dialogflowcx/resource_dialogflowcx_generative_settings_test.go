package dialogflowcx_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
)

func TestAccDialogflowCXGenerativeSettings_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":          envvar.GetTestOrgFromEnv(t),
		"billing_account": envvar.GetTestBillingAccountFromEnv(t),
		"random_suffix":   acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDialogflowCXGenerativeSettings_full(context),
			},
			{
				ResourceName:            "google_dialogflow_cx_generative_settings.my_generative_settings",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"fallback_settings.0.prompt_templates"},
			},
			{
				Config: testAccDialogflowCXGenerativeSettings_update(context),
			},
			{
				ResourceName:            "google_dialogflow_cx_generative_settings.my_generative_settings",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"fallback_settings.0.prompt_templates"},
			},
		},
	})
}

func testAccDialogflowCXGenerativeSettings_full(context map[string]interface{}) string {
	return acctest.Nprintf(`
  resource "google_dialogflow_cx_agent" "agent" {
    display_name               = "tf-test-%{random_suffix}update"
    location                   = "global"
    default_language_code      = "en"
    time_zone                  = "America/New_York"
    description                = "Example description."
  }

  resource "google_dialogflow_cx_generative_settings" "my_generative_settings" {
    parent       = google_dialogflow_cx_agent.agent.id

    fallback_settings {
      selected_prompt = "example prompt"
      prompt_templates {
        display_name = "example prompt"
        prompt_text = "example prompt text"
        frozen = false
      }
    }

    generative_safety_settings {
      default_banned_phrase_match_strategy = "PARTIAL_MATCH"
      banned_phrases {
        text = "example text"
        language_code = "en"
      }
    }

    knowledge_connector_settings {
      business = "example business"
      agent = "example agent"
      agent_identity = "virtual agent"
      business_description = "a family company selling freshly roasted coffee beans"
      agent_scope = "Example company website"
      disable_data_store_fallback = false
    }

    language_code = "en"

    llm_model_settings {
      model = "gemini-2.0-flash-001"
      prompt_text = "example prompt text"
    }
  }
`, context)
}

func testAccDialogflowCXGenerativeSettings_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
  resource "google_dialogflow_cx_agent" "agent" {
      display_name               = "tf-test-%{random_suffix}update"
    location                   = "global"
    default_language_code      = "en"
    time_zone                  = "America/New_York"
    description                = "Example description."
  }

  resource "google_dialogflow_cx_generative_settings" "my_generative_settings" {
    parent       = google_dialogflow_cx_agent.agent.id

    knowledge_connector_settings {
      business = "updated business"
      agent = "updated agent"
    }

    fallback_settings {
      selected_prompt = "example prompt"
      prompt_templates {
        display_name = "example prompt"
        prompt_text = "example prompt text"
        frozen = false
      }
    }

    language_code = "en"
  }
`, context)
}
